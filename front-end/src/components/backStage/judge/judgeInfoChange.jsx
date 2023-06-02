import SingleInfo from "./singleInfo"
import Title from "../title"
import { useEffect, useState } from "react"
import ChangedInfoDetail from "./judgeInfoChange/ChangedInfoDetail";
import useAxios from "../../../utils/useAxios";
function JudgeInfoChange() {
  const axios = useAxios();
  const [list, setList] = useState([]);

  const [company, setCompany] = useState(null);

  async function getList() {
    const res = await axios.get("/manage/company/update-info/list");
    const list = res.data.data.CompaniesUpdateInfoList;
    setList(list)

  }

  useEffect(() => {
    getList();
  },[])

  async function pass(info) {
    const res = await axios.post('/manage/company/info',{
      company_accounts: [info.account]
    })
    setList(list.filter((company) => company.account !== info.account))
    setCompany(null);
  }
  
  async function reject(info) {
    const res = await axios.post('/manage/company/reject-info',{
      company_accounts: [info.account]
    })
    setList(list.filter((company) => company.account !== info.account))
    setCompany(null)
  }


  return <div className="rounded-3xl bg-gray-100 p-[2rem] h-[calc(100vh_-_4rem)]">
    <Title title="审批信息修改"></Title>
    <div className="flex flex-col gap-[1rem] mt-[2rem] h-[calc(100%_-_4rem)] scrollbar-blue overflow-y-auto pr-3">
      {list.map((info) => <SingleInfo key={info.account} info={info} handle={() => setCompany(info)}></SingleInfo>)}
    </div>
    {company?<ChangedInfoDetail info={company} close={() => setCompany(null)} pass={() => pass(company)} reject={() => reject(company)}></ChangedInfoDetail>:null}
  </div>
}

export default JudgeInfoChange