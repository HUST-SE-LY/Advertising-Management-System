import { useState } from "react";
import Title from "../backStage/title";

function MyAd() {
  const [list, setList] = useState([
    {
      id: 1,
      state: "已过期",
      start: "2023-1-1",
      expire: "2023-3-26",
    },
    {
      id: 1,
      state: "正常",
      start: "2023-4-1",
      expire: "2023-6-1",
    },
    {
      id: 2,
      state: "待展示",
      start: "2023-5-1",
      expire: "2023-6-1",
    },
  ]);

  return (
    <div className="row-span-1 col-span-1 bg-gray-50 rounded-[2rem] p-[2rem]">
      <Title title="持有广告位"></Title>
      <div className="mt-[1rem] flex flex-col gap-[1rem] h-[calc(100%_-_3rem)] overflow-y-auto scrollbar-blue pr-3">
        <div className="flex w-full">
          <p className="w-1/4 text-center">广告位</p>
          <p className="w-1/4 text-center">状态</p>
          <p className="w-1/4 text-center">开始时间</p>
          <p className="w-1/4 text-center">结束时间</p>
        </div>
        {
          list.map((info) => <div className="flex w-full bg-blue-100 py-[1rem] rounded-xl">
          <p className="w-1/4 text-center">{info.id}</p>
          <p className={`w-1/4 text-center ${info.state === '正常' ? 'text-green-500' : info.state === '待展示' ? 'text-blue-500' : 'text-red-500'}`}>{info.state}</p>
          <p className="w-1/4 text-center">{info.start}</p>
          <p className="w-1/4 text-center">{info.expire}</p>
        </div>)
        }
      </div>
    </div>
  );
}

export default MyAd;
