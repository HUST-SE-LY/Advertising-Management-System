import Title from "../title";
import LongButton from "../../longButton";
import { useNavigate } from "react-router-dom";
function Judge(props) {

  const navigate = useNavigate()
  return (
    <div className="bg-blue-50 rounded-3xl p-[2rem] w-[15rem] h-[15rem] animate-floatin">
      <p className="text-center text-gray-700">{props.title}</p>
      <p className="text-[4rem] text-center font-bold text-gray-700/70">{props.num}</p>
      {
        parseInt(props.num)?<LongButton content="前往处理" handle={() => {navigate("/back-stage/judge")}}></LongButton>:null
      }
      
    </div>
  );
}

export default Judge;
