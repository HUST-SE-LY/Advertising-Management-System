function LongButton(props) {

  return (<button className={`w-full p-1 h-fit text-center py-[0.5rem] rounded-lg mt-[1rem] transition-all hover:shadow-lg text-white ${props.color==='red'? 'bg-red-600/80 hover:bg-red-600 hover:shadow-red-300/5':'hover:shadow-blue-300/5 bg-blue-600/80 hover:bg-blue-600'} `} onClick={props.handle}>{props.content}</button>)
}

export default LongButton;