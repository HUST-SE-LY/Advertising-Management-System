import axios from "axios";

function useAxios() {
  const instance = axios.create({
    timeout: 3000000,
    baseURL: "https://aivwe.top/",
  });

  instance.interceptors.request.use((config) => {
    const token = localStorage.getItem("token");
    if (token) {
      config.headers.Authorization = `${token}`;
      console.log(config);
    }
    return config;
  });

  instance.interceptors.response.use((config) => {
    console.log(config);
    if (config.data.data) {
      config.data.data = JSON.parse(config.data.data);
    }
    return config;
  });

  return instance;
}

export default useAxios;
