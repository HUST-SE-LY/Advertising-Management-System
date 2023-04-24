import { useState } from "react"

function NoBorderInput(props) {

  const [value, setValue] = useState(props.value?props.value:"")

  function handle(e) {
    if(e.keyCode === 13) {
      props.handle();
    }
  }

  return <div className="my-[2rem] relative">
    <input type={props.type?props.type:'text'} className="w-full border-b border-gray-600 transition-all peer focus:border-blue-600 focus:text-blue-600 outline-none" value={value} onChange={(e) => setValue(e.target.value) } onKeyUp={(e) => handle(e)} />
    <p className={` pointer-events-none peer-focus:text-blue-600 text-gray-600 absolute top-0 peer-focus:text-sm peer-focus:top-[-1rem] transition-all ${value?'top-[-1rem] text-sm':''}`}>{props.children}</p>
  </div>
}

export default NoBorderInput