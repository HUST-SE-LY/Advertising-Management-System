import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Title from "../title";
import useAxios from "../../../utils/useAxios";

function CompanyList() {
  const axios = useAxios()
  const [list, setList] = useState([]);
  async function getCompanyList() {
    const res = await axios.get("/manage/company/list");
    setList(res.data.data.company_infos)
  }
  useEffect(() => {
    getCompanyList()
  },[])
  return (
    <div className="bg-blue-50 rounded-3xl h-full row-span-2 col-span-3 p-[2rem] animate-floatin">
      <Title title={"客户总览"}></Title>
      <div>
        <div className="flex p-[1rem]">
          <p className="w-[10rem]">名称</p>
          <p className="w-[10rem]">负责人</p>
          <p>负责人电话</p>
        </div>
      </div>
      <div className="flex flex-col gap-[1rem] h-[calc(100%_-_5rem)] overflow-y-auto pl-3 pr-3 scrollbar-blue">
        {list.map((info) => (
          <SingleCompany key={info.account} info={info}></SingleCompany>
        ))}
      </div>
    </div>
  );
}

function SingleCompany(props) {
  const navigate = useNavigate()
  return (
    <div className="flex bg-blue-100 rounded-xl items-center p-[1rem] pr-[2rem] shadow-lg shadow-blue-600/10">
      <p className="w-[10rem]">{props.info.name}</p>
      <p className="w-[10rem]">{props.info.manager_name}</p>
      <p>{props.info.manager_tel}</p>
      <button onClick={() => navigate(`/back-stage/companies/${props.info.account}`)} className="px-[1rem] py-[0.5rem] rounded-2xl bg-blue-500 text-white hover:shadow-lg hover:shadow-blue-600/20 hover:bg-blue-600 transition-all ml-[auto]">详情</button>
    </div>
  );
}

export default CompanyList;
