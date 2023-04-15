import { useState } from "react";

function PhoneModel(props) {
  const [list, setList] = useState([null,null,null,null,null,null,null,null,null,null])

  return <div className="justify-self-end bg-white w-full h-full border-2 border-blue-200 rounded-[2rem] min-w-[15rem] max-w-[20rem] animate-fadein overflow-hidden shadow-xl shadow-blue-600/10">
    <div className="h-[2rem] w-full rounded-[2rem_2rem_0_0] flex justify-center items-center">
      <div className="w-full">
        <p className="ml-[2rem]">5:20</p>
      </div>
      <div className="flex-shrink-0 w-[0.5rem] h-[0.5rem] bg-gray-700 rounded-full"></div>
      <div className="w-full"></div>
    </div>
    <div className=" h-[calc(100%_-_2rem)] w-full  bg-cover bg-center rounded-[0_0_2rem_2rem] relative overflow-y-auto scrollbar-blue">
      <div className="text-sm absolute top-0 left-0 text-center w-full h-[4rem] bg-blue-200/50 rounded-[0_0_1rem_1rem] leading-[4rem]">1号广告位</div>
      <div className="absolute left-0 top-[6rem] text-center w-full h-[10rem]">
        <p className="text-2xl tracking-[0.25rem] font-bold text-blue-600/70">Coogle</p>
        <input className="w-4/5 outline-none border border-blue-200 h-[2rem] leading-[2rem] rounded-[1rem] pl-[1rem] mt-[1rem]"/>
        <div className="text-center mx-auto mt-[8rem] rounded-xl w-[12rem] h-[4rem] bg-blue-200/50 leading-[4rem]">10号广告位</div>
      </div>
      <div className="flex absolute w-[12rem] mx-auto h-[6rem] top-[12rem] left-[calc(50%_-_6rem)] flex-wrap">
          {list.slice(2, 10).map((info,index) => <div className="border border-blue-200 box-border w-[3rem] h-[3rem] text-center leading-[3rem]">{info?"":index + 2}</div>)}
      </div>
      
    </div>
  </div>
}

export default PhoneModel;