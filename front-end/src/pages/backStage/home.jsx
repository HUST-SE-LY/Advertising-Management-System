import CompanyList from "../../components/backStage/home/companyList"
import Ads from "../../components/backStage/home/ads"
import AdsForSale from "../../components/backStage/home/adsForSale"
import Judge from "../../components/backStage/home/judge"
function BackStageHome() {
  return <div className="grid grid-cols-3 gap-[2rem]">
    <CompanyList></CompanyList>
    <Ads></Ads>
    <AdsForSale></AdsForSale>
    <Judge></Judge>
  </div>
}

export default BackStageHome