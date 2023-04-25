import { useState } from "react";
import Title from "../backStage/title";

function AdInfo(props) {
  const [adList, setAdList] = useState([
    {
      id: 0,
      state: 'not free',
      price: 100,
    }
    ,
    {
      id: 1,
      state: 'free',
      price: 10,
    }
    ,
    {
      id: 2,
      state: 'not free',
      price: 10,
    }
    ,
    {
      id: 3,
      state: 'free',
      price: 10,
    }
    ,
    {
      id: 4,
      state: 'free',
      price: 10,
    }
    ,
    {
      id: 5,
      state: 'free',
      price: 10,
    }
    ,
    {
      id: 6,
      state: 'free',
      price: 10,
    }
    ,
    {
      id: 7,
      state: 'free',
      price: 10,
    }
    ,
    {
      id: 8,
      state: 'not free',
      price: 10,
    }
    ,
    {
      id: 9,
      state: 'free',
      price: 100,
    }
  ])

  return <div className="row-span-1 col-span-1 bg-gray-50 rounded-[2rem] p-[2rem] animate-listItemIn">
    <Title title="广告位信息"></Title>
    <div className="mt-[1rem] flex flex-col gap-[1rem] h-[calc(100%_-_3rem)] overflow-y-auto scrollbar-blue pr-3">
      {
        adList.map((info) => <div className="flex gap-[1rem] bg-blue-100 items-center py-[0.5rem] px-[1rem] rounded-2xl">
          <p className="w-[7rem]">{info.id + 1}号广告位</p>
          <p className={`w-[4rem] ${info.state === 'free'?'text-blue-500':'text-red-500'}`}>{info.state === 'free'?'空闲':'已售'}</p>
          <p className="">{info.price}元/天</p>
          <button className="ml-[auto] px-[1rem] py-[0.5rem] bg-blue-300 rounded-xl transition-all hover:bg-blue-600 hover:shadow-xl hover:shadow-blue-600/10 hover:text-white">{info.state === 'free'?'购买':'预定'}</button>
        </div>)
      }
    </div>
  </div>
}

export default AdInfo;