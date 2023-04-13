import { useState } from "react";
import SearchBar from "../../components/searchBar";
import CompanyDetail from "../../components/backStage/company/companyDetail";
import { Link, NavLink, Outlet } from "react-router-dom";

function BackStageCompany() {
  const [list, setList] = useState([
    {
      account: "1",
      name: "A",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },
    {
      account: "2",
      name: "B",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },{
      account: "3",
      name: "C",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },{
      account: "4",
      name: "D",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },{
      account: "5",
      name: "D",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },{
      account: "6",
      name: "D",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },{
      account: "7",
      name: "D",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },{
      account: "8",
      name: "D",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },{
      account: "9",
      name: "D",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },{
      account: "10",
      name: "D",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },{
      account: "11",
      name: "D",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    }
  ]);

  return (
    <div className="w-full h-[calc(100vh_-_4rem)] bg-gray-50 rounded-2xl grid grid-cols-[1fr_3.5fr]">
      <div className="bg-gray-100 rounded-2xl p-[2rem] h-[calc(100vh_-_4rem)]">
        <SearchBar searchKey="name" data={list} setInfo={setList}></SearchBar>
        <div className="mt-[2rem] flex flex-col gap-[1rem] h-[calc(100%_-_4.5rem)] pr-3 overflow-y-auto">
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
