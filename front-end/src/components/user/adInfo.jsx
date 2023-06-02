import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Title from "../backStage/title";
import useAxios from "../../utils/useAxios";

function AdInfo(props) {
  const navigate = useNavigate()
  const axios = useAxios();
  async function getAdList() {
    const res = await axios.get("/company/space");
    setAdList(res.data.data.space)
  }

  const [adList, setAdList] = useState([])
  useEffect(() => {
    getAdList()
  },[])

  return <div className="row-span-1 col-span-1 bg-gray-50 rounded-[2rem] p-[2rem] animate-listItemIn">
    <Title title="广告位信息"></Title>
    <div className="mt-[1rem] flex flex-col gap-[1rem] h-[calc(100%_-_3rem)] overflow-y-auto scrollbar-blue pr-3">
      {
        adList.map((info) => <div key={info.id} className="flex gap-[1rem] bg-blue-100 items-center py-[0.5rem] px-[1rem] rounded-2xl">
          <p className="w-[7rem]">{info.id}号广告位</p>
          <p className="">{info.price}元/天</p>
          <button onClick={() => navigate(`/buy/${info.id}`)} className="ml-[auto] px-[1rem] py-[0.5rem] bg-blue-300 rounded-xl transition-all hover:bg-blue-600 hover:shadow-xl hover:shadow-blue-600/10 hover:text-white">购买</button>
        </div>)
      }
    </div>
  </div>
}

export default AdInfo;