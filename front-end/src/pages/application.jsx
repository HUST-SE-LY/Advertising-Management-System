import { useState } from "react";

function Application(props) {

  const [list,setList] = useState([
    {
      url: 'https://github.com/',
      image: 'https://w.wallhaven.cc/full/g7/wallhaven-g77oz7.jpg'
    },
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
  ])

  return <div className="w-screen h-screen flex justify-center items-center">
    <div className="fixed top-0 left-0 w-full h-[20%] ">
      {
        list[0]?<a target="_blank" href={list[0].url}><img src={list[0].image} className="max-w-full max-h-full mx-[auto] rounded-xl shadow-2xl shadow-gray-200" alt="" /></a>:"广告位空闲"
      }
    </div>
    <div className="w-full flex h-fit flex-col gap-[2rem] items-center">
      <p className="text-[4rem] font-bold font-sans tracking-widest text-blue-600">Coogle</p>
      <input type="text"  className=" outline-none border-2 w-3/5 border-blue-400 px-[2rem] h-[4rem] rounded-[2rem] text-xl"/>
      <div className="grid grid-cols-4 grid-rows-2 gap-[1rem] relative">
        {
          list.slice(1, 9).map((info,index) => info?<div className="text-center  max-sm:w-[3rem] max-sm:h-[3rem] sm:w-[4rem] sm:h-[4rem] md:w-[6rem] md:h-[6rem] lg:w-[10rem] lg:h-[10rem] flex items-center justify-center"><a target="_blank" href={list[index + 1].url}><img src={list[index + 1].image} className="max-w-full max-h-full mx-[auto] rounded-xl shadow-2xl shadow-gray-200" alt="" /></a></div>:<div className="text-center max-sm:w-[3rem] max-sm:h-[3rem]  max-sm:text-sm sm:w-[4rem] sm:h-[4rem] sm:text-sm md:w-[6rem] md:h-[6rem] lg:w-[10rem] lg:h-[10rem] flex items-center justify-center border rounded-xl border-blue-200 shadow-xl shadow-blue-200/10">空闲广告位</div>)
        }
        
        <div className="flex items-center justify-center w-full h-[10rem] lg-height:h-[10rem] md-height:h-[6rem] md-height:bottom-[-8rem] border absolute bottom-[-15rem] rounded-xl border-blue-200 shadow-xl shadow-blue-200/10">空闲广告位</div>
      </div>
    </div>
  </div>
}

export default Application;