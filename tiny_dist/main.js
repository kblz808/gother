
async function readFileIntoBuffer(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => {
      resolve(reader.result);
    };
    reader.onerror = reject;
    reader.readAsArrayBuffer(file);
  });
}

async function init(){
  const go = new Go();
  const WASM_URL= 'main.wasm';
  let wasm;

  if ('instantiateStreaming' in WebAssembly) {
  	WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(function (obj) {
  		wasm = obj.instance;
  		go.run(wasm);
  	})
  } else {
  	fetch(WASM_URL).then(resp =>
  		resp.arrayBuffer()
  	).then(bytes =>
  		WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
  			wasm = obj.instance;
  			go.run(wasm);
  		})
  	)
  }
  
  document.getElementById('imageUpload').addEventListener('change', async function(event){
    const file = event.target.files[0];

    let buffer = await readFileIntoBuffer(file);
    buffer = new Uint8Array(buffer);
    
    decodeImage(buffer)
  })
  
}

init();

function displayImage(encodedImage) {
  const img = document.createElement("img");
  img.src = "data:image/jpeg;base64," + encodedImage;
  document.body.appendChild(img);
}
