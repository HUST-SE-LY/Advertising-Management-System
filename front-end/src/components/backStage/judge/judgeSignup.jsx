import { useEffect, useState } from "react";
import Title from "../title";
import SingleInfo from "./singleInfo";
import SignupInfo from './judgeSignup/signupInfo'
import useAxios from "../../../utils/useAxios";


function JudgeSignup() {
  const axios  = useAxios();
  const [list,setList] = useState([])
  async function getList() {
    const res = await axios.get("/manage/company/review");
    setList(res.data.data.company_infos)
  }
  useEffect(() => {
    getList()
  },[])
  const [company, setCompany] = useState(null);

  async function passCompany(info) {
    const res = await axios.post("/manage/company/register",{
      company_accounts: [info.account]
    })
    console.log(res);
    setList(list.filter((company) => company.account !== info.account))
    setCompany(null);
  }
  
  function rejectCompany(info) {
    setList(list.filter((company) => company.account !== info.account))
    setCompany(null)
  }

  return <div className="rounded-3xl bg-gray-100 p-[2rem] h-[calc(100vh_-_4rem)]">
    <Title title="审批注册">
    </Title>
    <div className="flex flex-col gap-[1rem] mt-[2rem] h-[calc(100%_-_4rem)] scrollbar-blue overflow-y-auto pr-3">
      {list.map((info) => <SingleInfo key={info.account} info={info} handle={() => setCompany(info)}></SingleInfo>)}
    </div>
    {company?<SignupInfo info={company} close={() => setCompany(null)} pass={() => passCompany(company)} reject={() => rejectCompany(company)}></SignupInfo>:null}
  </div>
}

export default JudgeSignup