import { useState } from "react";
import Title from "../title";
import AdInfo from "./judgeAd/AdInfo";
import SingleInfo from "./singleInfo";

function JudgeAd() {

  const [list, setList] = useState([{
    AdId: '0',
    account: '1',
    name: 'AAA',
    image: '/src/assets/testAd.jpg',
    url: 'https://github.com/HUST-SE-LY/Advertising-Management-System',
  },{
    AdId: '1',
    account: '1',
    name: 'AAA',
    image: '/src/assets/testAd.jpg',
    url: 'https://github.com/HUST-SE-LY/Advertising-Management-System',
  },{
    AdId: '2',
    account: '1',
    name: 'AAA',
    image: '/src/assets/testAd.jpg',
    url: 'https://github.com/HUST-SE-LY/Advertising-Management-System',
  },{
    AdId: '3',
    account: '1',
    name: 'AAA',
    image: '/src/assets/testAd.jpg',
    url: 'https://github.com/HUST-SE-LY/Advertising-Management-System',
  }])

  const [Ad, setAd] = useState(null)

  function passAd() {
    setList(list.filter(info => info.AdId !== Ad.AdId));
    setAd(null);
  }

  function rejectAd() {
    setList(list.filter(info => info.AdId !== Ad.AdId));
    setAd(null);
  }


  return <div className="rounded-3xl bg-gray-100 p-[2rem]">
    <Title title="审批广告申请"></Title>
    <div className="flex flex-col gap-[1rem] mt-[2rem]">
      {list.map((info) => <SingleInfo key={info.AdId} info={info} handle={() => setAd(info)}></SingleInfo>)}
    </div>
    {Ad?<AdInfo info={Ad} close={() => setAd(null)} pass={passAd} reject={rejectAd} ></AdInfo>:null}
  </div>

}

export default JudgeAd;