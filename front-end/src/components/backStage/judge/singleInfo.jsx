
function SingleInfo(props) {
  const name = props.info.name;
  return <div className="bg-white rounded-xl shadow-lg shadow-gray-300/20 h-fit p-[1rem] flex gap-[1rem] items-center animate-listItemIn">
    <p>{name}</p>
    <div className="ml-[auto] flex items-center gap-[1rem]">
      <button className="bg-blue-500 hover:bg-blue-600 hover:shadow-blue-600/30 shadow-xl shadow-blue-600/10 transition-all rounded-2xl w-[5rem] py-[10px] text-white overflow-hidden" onClick={props.handle}>详情</button>
    </div>
  </div>
}

export default SingleInfo