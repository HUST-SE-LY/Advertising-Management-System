import { useState } from "react";

function PcModel() {
  const [list, setList] = useState([
    {
      url: 'https://github.com/',
      image: 'https://w.wallhaven.cc/full/g7/wallhaven-g77oz7.jpg'
    },
    null,
    null,
    null,
    null,
    null,
    null,
    null,
    null,
    null,
  ]);

  return (
    <div className="border-2 relative border-blue-200 rounded-[2rem] shadow-xl shadow-blue-600/10 animate-fadein ">
      <div className="h-[2rem] w-full bg-white rounded-[2rem_2rem_0_0] flex gap-[0.5rem] items-center px-[2rem]">
        <div className="w-[1rem] h-[1rem] bg-red-400 rounded-full"></div>
        <div className="w-[1rem] h-[1rem] bg-yellow-400 rounded-full"></div>
        <div className="w-[1rem] h-[1rem] bg-green-400 rounded-full"></div>
      </div>
      <div className="h-[calc(100%_-_2rem)] flex  justify-center relative overflow-y-auto scrollbar-blue">
        <div className="text-center w-full mt-[4rem]">
          <p className="text-2xl tracking-[0.25rem] font-bold text-blue-600/70">
            Coogle
          </p>
          <input className="w-4/5 outline-none border border-blue-200 h-[2rem] leading-[2rem] rounded-[1rem] pl-[1rem] mt-[1rem]" />
          <div className="flex w-[16rem] mx-auto h-[8rem] flex-wrap mt-[1rem]">
            {list.slice(2, 10).map((info, index) => (
              <div className="border border-blue-200 box-border w-[4rem] h-[4rem] text-center leading-[4rem]">
                {info ? "" : index + 2}
              </div>
            ))}
          </div>
          <div className="text-center mx-auto mt-[1rem] rounded-xl w-[16rem] h-[4rem] bg-blue-200/50 leading-[4rem]">10号广告位</div>
        </div>
        <div className={`absolute top-0 left-0 text-center w-full h-[4rem] bg-blue-200/50 rounded-[0_0_1rem_1rem] leading-[4rem]`}>
          1号广告位
        </div>
      </div>
    </div>
  );
}

export default PcModel;
