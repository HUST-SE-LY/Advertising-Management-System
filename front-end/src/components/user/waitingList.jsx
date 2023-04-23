import { useState } from "react";
import Title from "../backStage/title";

function WaitingList() {
  const [list, setList] = useState([
    {
      time: "2023-1-6",
      theme: "投放广告申请",
      state: "处理中",
    },
    {
      time: "2023-2-6",
      theme: "修改信息申请",
      state: "已通过",
    },
    {
      time: "2023-1-6",
      theme: "投放广告申请",
      state: "被拒绝",
    },
  ]);

  return (
    <div className="row-span-1 col-span-1 bg-gray-50 rounded-[2rem] p-[2rem]">
      <Title title="申请列表"></Title>
      <div className="mt-[1rem] flex flex-col gap-[1rem] h-[calc(100%_-_3rem)] overflow-y-auto scrollbar-blue pr-3">
        <div className="flex w-full">
          <p className="w-1/3 text-center">时间</p>
          <p className="w-1/3 text-center">主题</p>
          <p className="w-1/3 text-center">状态</p>
        </div>
        {list.map((info) => (
          <div className="flex w-full bg-blue-100 py-[1rem] rounded-xl">
            <p className="w-1/3 text-center">{info.time}</p>
            <p className="w-1/3 text-center">{info.theme}</p>
            <p className={`w-1/3 text-center ${info.state === '已通过' ? "text-green-500" : info.state === '被拒绝'?'text-red-500':'text-blue-500'}`}>{info.state}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default WaitingList;
