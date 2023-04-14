import { useState } from "react";
import Title from "../title";
import AdInfo from "./judgeAd/AdInfo";
import SingleInfo from "./singleInfo";

function JudgeAd() {

  const [list, setList] = useState([{
    AdId: '0',
    position: "1",
    account: '1',
    name: 'AAA',
    image: '/src/assets/testAd.jpg',
    url: 'https://github.com/HUST-SE-LY/Advertising-Management-System',
  },{
    AdId: '1',
    account: '1',
    position: "1",
    name: 'AAA',
    image: '/src/assets/testAd.jpg',
    url: 'https://github.com/HUST-SE-LY/Advertising-Management-System',
  },{
    AdId: '2',
    account: '1',
    position: "1",
    name: 'AAA',
    image: '/src/assets/testAd.jpg',
    url: 'https://github.com/HUST-SE-LY/Advertising-Management-System',
  },{
    AdId: '3',
    account: '1',
    position: "1",
    name: 'AAA',
    image: '/src/assets/testAd.jpg',
    url: 'https://github.com/HUST-SE-LY/Advertising-Management-System',
  },{
    AdId: '4',
    account: '1',
    position: "1",
    name: 'AAA',
    image: '/src/assets/testAd.jpg',
    url: 'https://github.com/HUST-SE-LY/Advertising-Management-System',
  },{
    AdId: '5',
    account: '1',
    position: "1",
    name: 'AAA',
    image: '/src/assets/testAd.jpg',
    url: 'https://github.com/HUST-SE-LY/Advertising-Management-System',
  },{
    AdId: '6',
    account: '1',
    position: "1",
    name: 'AAA',
    image: '/src/assets/testAd.jpg',
    url: 'https://github.com/HUST-SE-LY/Advertising-Management-System',
  },{
    AdId: '3',
    account: '1',
    position: "1",
    name: 'AAA',
    image: '/src/assets/testAd.jpg',
    url: 'https://github.com/HUST-SE-LY/Advertising-Management-System',
  },{
    AdId: '7',
    account: '1',
    position: "1",
    name: 'AAA',
    image: '/src/assets/testAd.jpg',
    url: 'https://github.com/HUST-SE-LY/Advertising-Management-System',
  },{
    AdId: '8',
    account: '1',
    position: "1",
    name: 'AAA',
    image: '/src/assets/testAd.jpg',
    url: 'https://github.com/HUST-SE-LY/Advertising-Management-System',
  },{
    AdId: '9',
    account: '1',
    position: "1",
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


  return <div className="relative rounded-3xl bg-gray-100 p-[2rem] h-[calc(100vh_-_4rem)]">
    <Title title="审批广告申请"></Title>
    <div className="flex flex-col gap-[1rem] mt-[2rem] h-[calc(100%_-_4rem)] overflow-y-auto pr-3">
      {list.map((info) => <SingleInfo key={info.AdId} info={info} handle={() => setAd(info)}></SingleInfo>)}
    </div>
    {Ad?<AdInfo info={Ad} close={() => setAd(null)} pass={passAd} reject={rejectAd} ></AdInfo>:null}
  </div>

}

export default JudgeAd;