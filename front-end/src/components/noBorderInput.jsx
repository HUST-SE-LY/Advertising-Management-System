import { useEffect, useState } from "react"

function NoBorderInput(props) {
  const originValue = props.value;
  const [value, setValue] = useState(originValue);
  useEffect(() => {
    setValue(originValue);
  },[originValue])
  function changeInfo(e) {
    setValue(e.target.value);
    props.setInfo(e.target.value);
  }
  function handle(e) {
    if(e.keyCode === 13) {
      props.handle();
    }
  }

  return <div className="my-[2rem] relative">
    <input type={props.type?props.type:'text'} className="w-full border-b border-gray-600 transition-all peer focus:border-blue-600 focus:text-blue-600 outline-none" value={value} onChange={changeInfo} onKeyUp={(e) => handle(e)} />
    <p className={` pointer-events-none peer-focus:text-blue-600 text-gray-600 absolute top-0 peer-focus:text-sm peer-focus:top-[-1rem] transition-all ${value?'top-[-1rem] text-sm':''}`}>{props.children}</p>
  </div>
}

export default NoBorderInput