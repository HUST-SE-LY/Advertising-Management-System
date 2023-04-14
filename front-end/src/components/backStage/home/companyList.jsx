import { useState } from "react";
import Title from "../title";

function CompanyList() {
  const [list, setList] = useState([
    {
      account: "1",
      name: "company",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },
    {
      account: "2",
      name: "company2",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },
    {
      account: "3",
      name: "company3",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },
    {
      account: "4",
      name: "company4",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },
    {
      account: "5",
      name: "company4",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },
    {
      account: "6",
      name: "company4",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },
    {
      account: "7",
      name: "company4",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },
    {
      account: "8",
      name: "company4",
      address: "china",
      managerName: "me",
      managerTel: "123456",
      businessLicenseNumber: "1111111111",
    },
  ]);
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
      <div className="flex flex-col gap-[1rem] h-[calc(100%_-_5rem)] overflow-y-auto">
        {list.map((info) => (
          <SingleCompany key={info.account} info={info}></SingleCompany>
        ))}
      </div>
    </div>
  );
}

function SingleCompany(props) {
  return (
    <div className="flex bg-blue-100 rounded-xl items-center p-[1rem] pr-[2rem] shadow-lg shadow-blue-600/10">
      <p className="w-[10rem]">{props.info.name}</p>
      <p className="w-[10rem]">{props.info.managerName}</p>
      <p>{props.info.managerTel}</p>
      <button className="px-[1rem] py-[0.5rem] rounded-2xl bg-blue-500 text-white hover:shadow-lg hover:shadow-blue-600/20 hover:bg-blue-600 transition-all ml-[auto]">详情</button>
    </div>
  );
}

export default CompanyList;
