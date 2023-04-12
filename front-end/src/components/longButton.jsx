function LongButton(props) {

  return (<button className="w-full p-1 h-fit text-center py-[0.5rem] rounded-lg mt-[1rem] transition-all hover:shadow-lg hover:shadow-blue-300/50 text-white bg-blue-600/80 hover:bg-blue-600" onClick={props.handle}>{props.content}</button>)
}

export default LongButton;