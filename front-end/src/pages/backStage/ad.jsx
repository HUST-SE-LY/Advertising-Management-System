import PhoneModel from "../../components/backStage/ad/phoneModel";
import PcModel from "../../components/backStage/ad/pcModel";
import Title from "../../components/backStage/title";
import { NavLink, Outlet } from "react-router-dom";

function BackStageAd() {

  const idList = [1,2,3,4,5,6,7,8,9,10];

  return (
    <div className="w-full h-[calc(100vh_-_4rem)] bg-gray-50 rounded-2xl grid grid-cols-2 p-[1rem]">
      <div className="pr-[1rem] flex flex-col gap-[1rem] h-[calc(100vh_-_6rem)]">
        <div className="flex gap-[1rem] items-center">
          <div className="flex-shrink-0">广告位编号：</div>
          <div className="w-full overflow-x-auto scrollbar-blue rounded-xl flex gap-[1rem] p-[1rem] bg-blue-100">
            {
              idList.map((id) => <NavLink key={id} to={`/back-stage/ads/${id}`} className={({isActive}) => isActive?'rounded-xl px-[2rem] py-[0.5rem] transition-all bg-blue-600 text-white':'rounded-xl transition-all bg-blue-200 px-[2rem] py-[0.5rem]'}>{id}</NavLink>)
            }
          </div>
        </div>
        <Outlet></Outlet>
      </div>
      <div className="grid grid-rows-2 gap-[1rem] h-[calc(100vh_-_6rem)]">
        <div className="grid grid-cols-2 ">
          <div className="p-[2rem] flex flex-col gap-[1rem] text-lg justify-center">
            <p className="text-right text-2xl text-gray-700 font-bold">广告预览</p>
            <p className="text-right">下为PC端预览</p>
            <p className="text-right">右为移动端预览</p>
          </div>
          <PhoneModel></PhoneModel>
        </div>
        <PcModel></PcModel>
      </div>
    </div>
  );
}

export default BackStageAd;
