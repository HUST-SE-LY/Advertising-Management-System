import CompanyList from "../../components/backStage/home/companyList";
import AdsForSale from "../../components/backStage/home/adsForSale";
import Judge from "../../components/backStage/home/judge";
import LongButton from "../../components/longButton";
import AddManager from "../../components/backStage/home/addManager";
import useAxios from "../../utils/useAxios";

import { Link, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
function BackStageHome() {
  const axios = useAxios();
  const navigate = useNavigate()
  const [showAddManager, setShowAddManager] = useState(0);
  const [registerNum, setRegisterNum] = useState(0);
  const [modifyNum, setModifyNum] = useState(0);
  const [AdNum, setAdNum] = useState();
  async function checkLogin() {
    try {
      const res = await axios.get("/admin/check_login");
    } catch (err) {
      navigate("/");
    }
  }
  async function logout() {
    const res = await axios.post("/admin/logout");
    localStorage.removeItem("token");
    navigate("/")
  }
  async function getCompanyCount() {
    const res = await axios.get("/manage/company/review_count");
    setRegisterNum(res.data.data.count)
  }
  async function getInfoCount() {
    const res = await axios.get("/manage/company/info_review_count");
    setModifyNum(res.data.data.count)
  }
  async function getAdCount() {
    const res = await axios.get("/manage/advertisement/count");
    console.log(res)
    setAdNum(res.data.data.adcount)
  }
  useEffect(() => {
    checkLogin();
    getCompanyCount();
    getInfoCount();
    getAdCount()
  }, []);

  return (
    <div className="bg-gray-50 h-[calc(100vh_-_4rem)] rounded-3xl">
      <div className="w-full h-fit px-[2rem] py-[1rem] flex gap-[2rem] bg-blue-200 rounded-3xl shadow-xl shadow-blue-300/20 items-center">
        <p className="text-2xl font-bold tracking-widest text-gray-700 ">
          欢迎回来!
        </p>
        <div className="ml-[auto] flex gap-[2rem]">
          <Link to="/">
            <button className="bg-blue-500 px-[3rem] py-[0.5rem] rounded-2xl text-white transition-all hover:shadow-lg hover:shadow-blue-600/30 hover:bg-blue-600">
              回首页
            </button>
          </Link>
          <Link to="/app">
            <button className="bg-blue-500 px-[3rem] py-[0.5rem] rounded-2xl text-white transition-all hover:shadow-lg hover:shadow-blue-600/30 hover:bg-blue-600">
              前往应用
            </button>
          </Link>
        </div>
      </div>

      <div className="grid h-[calc(100%_-_4.5rem)] grid-cols-[15rem_15rem_15rem_1fr] grid-rows-[15rem_3fr_1fr] gap-[2rem] p-[2rem] ">
        <Judge title="待审批注册" num={registerNum}></Judge>
        <Judge title="待审批修改" num={modifyNum}></Judge>
        <Judge title="待审批广告" num={AdNum}></Judge>
        <AdsForSale></AdsForSale>
        <CompanyList></CompanyList>
        <div>
          <LongButton
            content="添加管理员"
            handle={() => setShowAddManager(true)}
          ></LongButton>
          <LongButton content="退出登录" color="red" handle={logout}></LongButton>
        </div>
      </div>
      {showAddManager ? (
        <AddManager close={() => setShowAddManager(false)}></AddManager>
      ) : null}
    </div>
  );
}

export default BackStageHome;
