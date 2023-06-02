import { useEffect, useState } from "react";
import Title from "../title";
import AdInfo from "./judgeAd/AdInfo";
import SingleInfo from "./singleInfo";
import useAxios from "../../../utils/useAxios";

function JudgeAd() {

  const axios = useAxios();

  const [list, setList] = useState([])

  async function getList() {
    const res = await axios.get('/manage/advertisement/review');
    console.log(res)
    setList(res.data.data.advertisement_infos)
  }

  const [Ad, setAd] = useState(null)

  async function passAd(info) {
    const res = await axios.post("/manage/advertisement/allow",{
      pending_numbers:[parseInt(info.id)]
    })
    console.log(res)
    setList(list.filter(info => info.id !== Ad.id));
    setAd(null);
  }

  async function rejectAd(info) {
    console.log(info)
    const res = await axios.post("/manage/advertisement/reject",{
      reject_numbers:[parseInt(info.id)]
    })
    setList(list.filter(info => info.id !== Ad.id));
    setAd(null);
  }

  useEffect(() => {
    getList();
  },[])


  return <div className="relative rounded-3xl bg-gray-100 p-[2rem] h-[calc(100vh_-_4rem)]">
    <Title title="审批广告申请"></Title>
    <div className="scrollbar-blue flex flex-col gap-[1rem] mt-[2rem] h-[calc(100%_-_4rem)] scrollbar-blue overflow-y-auto pr-3">
      {list.map((info) => <SingleInfo key={info.id} info={info} handle={() => setAd(info)}></SingleInfo>)}
    </div>
    {Ad?<AdInfo info={Ad} close={() => setAd(null)} pass={()=>passAd(Ad)} reject={()=>rejectAd(Ad)} ></AdInfo>:null}
  </div>

}

export default JudgeAd;