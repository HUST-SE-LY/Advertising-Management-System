import { useState } from "react"

function SearchBar(props) {
  const [value, setValue] = useState("");
  const originalList = [...props.data]
  console.log(originalList)
  function search(e) {
    if(e.keyCode == 13) {
      if(!value) {
        console.log(originalList)
        props.setInfo(originalList);
        return ;
      }
      props.setInfo(props.data.filter((info) => {
        return info[props.searchKey].includes(value)
      }))
      console.log(originalList)
    }
  }

  return <input value={value} onChange={(e) => {setValue(e.target.value)}} onKeyUp={search} placeholder="搜索企业名称" className="w-full py-[0.5rem] px-[1rem] rounded-[2rem] outline-none bg-white"/>
}

export default SearchBar