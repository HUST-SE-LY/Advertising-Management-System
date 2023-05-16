import axios from "axios";

function useAxios() {
  const instance = axios.create({
    timeout: 3000000,
    baseURL: "https://aivwe.top/"
  });

  instance.interceptors.request.use((config) => {
    const token = localStorage.getItem("token");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  });

  instance.interceptors.response.use((config) => {
    console.log(config);
    return config;
  })

  return instance;
}


export default useAxios;