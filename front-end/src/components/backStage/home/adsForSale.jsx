import { useState } from "react";
import Title from "../title";

function AdsForSale() {
  const [list, setList] = useState([{
    id: 0,
    state: 'free',
    price: '100',
  },{
    id: 1,
    state: 'saled',
    price: '200',
  },{
    id: 2,
    state: 'saled',
    price: '200',
  },{
    id: 3,
    state: 'saled',
    price: '200',
  },{
    id: 4,
    state: 'saled',
    price: '200',
  }])


  return (
    <div className="bg-blue-50 w-full row-span-2 rounded-3xl p-[2rem] animate-floatin">
      <Title title={"广告位总览"}></Title>
      <div className="flex flex-col gap-[1rem] mt-[1rem] h-[calc(100%_-_3rem)] overflow-y-auto">
        {
        list.map((info) => <SingleAds key={info.id} info={info}></SingleAds> )
      }
      </div>
      
    </div>
  );
}

function SingleAds(props) {
  return <div className="bg-blue-100 w-full flex gap-[2rem] p-[1rem] rounded-xl items-center animate-listItemIn shadow-lg shadow-blue-600/10">
    <p>{props.info.id}号广告栏</p>
    <p className={`${props.info.state === 'free'?'text-blue-600':'text-red-400'}`}>{props.info.state === 'free'?"空闲":'已售'}</p>
    <p>{props.info.price}元/天</p>
    <button className="px-[1rem] py-[0.5rem] bg-blue-500 text-white rounded-2xl transition-all hover:bg-blue-600 hover:shadow-lg hover:shadow-blue-600/20 ml-[auto]">管理</button>
  </div>
}

export default AdsForSale;
