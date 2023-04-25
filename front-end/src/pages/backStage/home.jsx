import CompanyList from "../../components/backStage/home/companyList"
import AdsForSale from "../../components/backStage/home/adsForSale"
import Judge from "../../components/backStage/home/judge"
import LongButton from "../../components/longButton"
import AddManager from "../../components/backStage/home/addManager"

import { Link } from "react-router-dom"
import { useState } from "react"
function BackStageHome() {
  const [showAddManager, setShowAddManager] = useState();
  return <div className="bg-gray-50 h-[calc(100vh_-_4rem)] rounded-3xl">
    <div className="w-full h-fit px-[2rem] py-[1rem] flex gap-[2rem] bg-blue-200 rounded-3xl shadow-xl shadow-blue-300/20 items-center">
      <p className="text-2xl font-bold tracking-widest text-gray-700 ">欢迎回来!</p>
      <div className="ml-[auto] flex gap-[2rem]">
        <Link to="/">
          <button className="bg-blue-500 px-[3rem] py-[0.5rem] rounded-2xl text-white transition-all hover:shadow-lg hover:shadow-blue-600/30 hover:bg-blue-600">回首页</button>
        </Link>
        <Link to="/app">
          <button className="bg-blue-500 px-[3rem] py-[0.5rem] rounded-2xl text-white transition-all hover:shadow-lg hover:shadow-blue-600/30 hover:bg-blue-600">前往应用</button>
        </Link>
      </div>
    </div>

    <div className="grid h-[calc(100%_-_4.5rem)] grid-cols-[15rem_15rem_15rem_1fr] grid-rows-[15rem_3fr_1fr] gap-[2rem] p-[2rem] ">
      <Judge title="待审批注册" num="0"></Judge>
      <Judge title="待审批修改" num="3"></Judge>
      <Judge title="待审批广告" num="5"></Judge>
      <AdsForSale></AdsForSale>
      <CompanyList></CompanyList>
      <div>
      <LongButton content="添加管理员" handle={() => setShowAddManager(true)}></LongButton>
        <LongButton content="退出登录" color="red"></LongButton>
      </div>
    </div>
    {
      showAddManager?<AddManager close={() => setShowAddManager(false)}></AddManager>:null
    }
  </div>
}

export default BackStageHome