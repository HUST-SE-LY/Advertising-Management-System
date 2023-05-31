import { useState, useEffect } from "react";
import SearchBar from "../../components/searchBar";
import CompanyDetail from "../../components/backStage/company/companyDetail";
import { Link, NavLink, Outlet } from "react-router-dom";
import useAxios from "../../utils/useAxios";

function BackStageCompany() {
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
    <div className="w-full h-[calc(100vh_-_4rem)] bg-gray-50 rounded-2xl grid grid-cols-[1fr_3.5fr]">
      <div className="bg-gray-100 rounded-2xl p-[2rem] h-[calc(100vh_-_4rem)]">
        <SearchBar searchKey="name" data={list} setInfo={setList}></SearchBar>
        <div className="mt-[2rem] flex flex-col gap-[1rem] h-[calc(100%_-_4.5rem)] pl-3 pr-3 scrollbar-blue overflow-y-auto">
          {list.map((info) => (
            <NavLink key={info.account}
              to={`/back-stage/companies/${info.account}`}
              className={({ isActive }) =>
                isActive ? "bg-blue-600 shadow-blue-600/50 shadow-lg text-white block w-full p-[1rem] rounded-2xl transition-all animate-listItemIn" : "block w-full p-[1rem] rounded-2xl hover:bg-blue-300 hover:shadow-lg hover:shadow-blue-300/50 bg-white transition-all animate-listItemIn"
              }
            >
              <div>{info.name}</div>
            </NavLink>
          ))}
        </div>
      </div>
      <Outlet></Outlet>
    </div>
  );
}

export default BackStageCompany;
