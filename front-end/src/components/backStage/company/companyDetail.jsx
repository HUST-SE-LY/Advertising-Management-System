import { useEffect } from "react";
import { useState } from "react";
import { useParams } from "react-router-dom";
import BasicInfo from "./companyDetail/basicInfo";

function CompanyDetail(props) {
  const params = useParams();
  const [info, setInfo] = useState(null);
  async function getInfo() {
    switch (params.id) {
      case "1":
        setInfo({
          account: "1",
          name: "A",
          address: "china",
          managerName: "me",
          managerTel: "123456",
          businessLicenseNumber: "1111111111",
        });
        break;
      case "2":
        setInfo({
          account: "2",
          name: "B",
          address: "china",
          managerName: "me",
          managerTel: "123456",
          businessLicenseNumber: "1111111111",
        });
        break;
      case "3":
        setInfo({
          account: "3",
          name: "C",
          address: "china",
          managerName: "me",
          managerTel: "123456",
          businessLicenseNumber: "1111111111",
        });
        break;
      case "4":
        setInfo({
          account: "4",
          name: "D",
          address: "china",
          managerName: "me",
          managerTel: "123456",
          businessLicenseNumber: "1111111111",
        });
        break;
    }
  }

  useEffect(() => {
    getInfo();
  },[params]);

  return <div>
    {info?<div className="w-full h-[calc(100vh_-_4rem)] p-[2rem] grid grid-cols-2 grid-rows-2 gap-[2rem]">
      <BasicInfo info={info}></BasicInfo>
    </div>:null}
  </div>
}

export default CompanyDetail;
