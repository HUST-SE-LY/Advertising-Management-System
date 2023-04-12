import { useEffect } from "react";
import { useState } from "react";

function TypingWord(props) {
  const len = props.word.length;
  const list = props.word;
  const [showList, setShowList] = useState("");
  useEffect(() => {
    if(showList.length === props.word.length) return ;
    const timeout = setTimeout(() => {
      setShowList(showList + list[showList.length])
    },300)
    return () => clearTimeout(timeout)
  }, [showList])
  return <div className="p-0 h-[200px] text-[5rem] leading-[200px] text-center font-sans font-bold text-blue-600/80">{showList}<span className="w-px h-[5rem] bg-blue-700 inline-block animate-shine"></span></div>;
}

export default TypingWord;
