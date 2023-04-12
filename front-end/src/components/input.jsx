import { useState } from "react"

function Input(props) {
  
  const [info, setInfo] = useState("");
  function changeInfo(e) {
    setInfo(e.target.value);
    props.setInfo(e.target.value);
  }
  return (<div>
    <p className="text-sm mb-1">{props.title}</p>
    <input type={props.type?props.type:'text'} className="border border-gray-400 outline-none p-1 rounded-lg w-full focus:border-blue-600 transition-all"/>
  </div>)

}

export default Input