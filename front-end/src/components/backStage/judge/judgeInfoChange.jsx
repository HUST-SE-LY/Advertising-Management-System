import SingleInfo from "./singleInfo"
import Title from "../title"
import { useState } from "react"
import ChangedInfoDetail from "./judgeInfoChange/ChangedInfoDetail";
function JudgeInfoChange() {
  const [list, setList] = useState([
    {
      account: '123456',
      name: '华科集团',
      newName: '武大集团',
      previousAddress: 'China',
      newAddress: 'America',
      previousManagerName: 'cheems',
      newManagerName: 'cheems',
      previousManagerTel:'123456',
      newManagerTel: '123456',
      previousLicense: '000000000',
      newLicense:'000000000',
    },    {
      account: '1234567',
      name: '华科集团',
      newName: '武大集团',
      previousAddress: 'China',
      newAddress: 'America',
      previousManagerName: 'cheems',
      newManagerName: 'cheems',
      previousManagerTel:'123456',
      newManagerTel: '123456',
      previousLicense: '000000000',
      newLicense:'000000000',
    },    {
      account: '1234568',
      name: '华科集团',
      newName: '武大集团',
      previousAddress: 'China',
      newAddress: 'America',
      previousManagerName: 'cheems',
      newManagerName: 'cheems',
      previousManagerTel:'123456',
      newManagerTel: '123456',
      previousLicense: '000000000',
      newLicense:'000000000',
    },    {
      account: '1234569',
      name: '华科集团',
      newName: '武大集团',
      previousAddress: 'China',
      newAddress: 'America',
      previousManagerName: 'cheems',
      newManagerName: 'cheems',
      previousManagerTel:'123456',
      newManagerTel: '123456',
      previousLicense: '000000000',
      newLicense:'000000000',
    }
  ])

  const [company, setCompany] = useState(null);

  function passCompany(info) {
    setList(list.filter((company) => company.account !== info.account))
    setCompany(null);
  }
  
  function rejectCompany(info) {
    setList(list.filter((company) => company.account !== info.account))
    setCompany(null)
  }


  return <div className="rounded-3xl bg-gray-100 p-[2rem] h-[calc(100vh_-_4rem)]">
    <Title title="审批信息修改"></Title>
    <div className="flex flex-col gap-[1rem] mt-[2rem] h-[calc(100%_-_4rem)] overflow-y-auto pr-3">
      {list.map((info) => <SingleInfo key={info.account} info={info} handle={() => setCompany(info)}></SingleInfo>)}
    </div>
    {company?<ChangedInfoDetail info={company} close={() => setCompany(null)} pass={() => passCompany(company)} reject={() => rejectCompany(company)}></ChangedInfoDetail>:null}
  </div>
}

export default JudgeInfoChange