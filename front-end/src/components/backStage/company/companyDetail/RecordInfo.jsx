function RecordInfo(props) {
  return (
    <div className="bg-gray-100 rounded-xl p-[2rem] h-full animate-floatin col-span-2">
      <p className="ml-[1rem] text-2xl font-bold tracking-wide relative before:absolute before:left-[-10px] before:w-[5px] before:h-full before:top-0 before:bg-blue-400">
        消费记录
      </p>
      <div className="pr-[calc(5px_+_0.75rem)]">
        <div className="flex w-full p-[1rem] text-center">
          <p className="w-1/5">订单号</p>
          <p className="w-1/5">下单时间</p>
          <p className="w-1/5">消费金额</p>
          <p className="w-1/5">购买广告位编号</p>
          <p className="w-1/5">购买时长</p>
        </div>
      </div>
      
      <div className="flex flex-col h-[calc(100%_-_5rem)] overflow-y-auto scrollbar-blue pr-3 gap-[1rem]">
        {props.info.map((info) => (
          <SingleRecord info={info} key={info.id}></SingleRecord>
        ))}
      </div>
    </div>
  );
}

function SingleRecord(props) {
  return (
    <div className="flex w-full bg-blue-100 p-[1rem] rounded-xl text-center animate-listItemIn">
      <p className="w-1/5">{props.info.id}</p>
      <p className="w-1/5">{props.info.time}</p>
      <p className="w-1/5">{props.info.price}</p>
      <p className="w-1/5">{props.info.AdPosition}</p>
      <p className="w-1/5">{props.info.AdTime}</p>
    </div>
  );
}

export default RecordInfo;
