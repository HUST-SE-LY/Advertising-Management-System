import { useEffect } from "react";
import { useState } from "react";
import { useParams } from "react-router-dom";
import BasicInfo from "./companyDetail/basicInfo";
import AccountInfo from "./companyDetail/accountInfo";
import RecordInfo from "./companyDetail/RecordInfo";

function CompanyDetail(props) {
  const params = useParams();
  const [info, setInfo] = useState(null);
  const [account, setAccount] = useState(null);
  const [record, setRecord] = useState(null);
  async function getInfo() {

  }

  async function getAccount() {
    setAccount({
      money: "114514",
      Ad: [
        {
          position: "1",
          leftTime: "10",
        },
      ],
    });
  }

  async function getRecord() {
    setRecord([
      {
        id: 0,
        time: "2023-3-25",
        price: "1000",
        AdPosition: "1",
        AdTime: "10",
      },
      {
        id: 1,
        time: "2023-3-23",
        price: "1000",
        AdPosition: "1",
        AdTime: "10",
      },
      {
        id: 2,
        time: "2023-3-26",
        price: "1000",
        AdPosition: "1",
        AdTime: "10",
      },{
        id: 3,
        time: "2023-3-26",
        price: "1000",
        AdPosition: "1",
        AdTime: "10",
      },{
        id: 4,
        time: "2023-3-26",
        price: "1000",
        AdPosition: "1",
        AdTime: "10",
      },
    ]);
  }

  useEffect(() => {
    getInfo();
    getAccount();
    getRecord();
  }, [params]);

  return (
    <div>
      {info&&account&&record ? (
        <div className="w-full h-[calc(100vh_-_4rem)] p-[2rem] grid grid-cols-2 grid-rows-2 gap-[2rem]">
          <BasicInfo info={info}></BasicInfo>
          <AccountInfo info={account}></AccountInfo>
          <RecordInfo info={record}></RecordInfo>
        </div>
      ) : null}
    </div>
  );
}

export default CompanyDetail;
