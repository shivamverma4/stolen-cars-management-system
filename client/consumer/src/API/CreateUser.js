import fetchData from "./../Utils/FetchData";
import Constants from "./../Constants/Constants";
import Url from "./../Config/Config";
import history from "./../Routes/history";

const createUser = async (userType, userDetails) => {
  const options = {
    method: Constants.httpMethods.POST,
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(userDetails)
  };
  try {
    const response = await fetchData(
      Url.apiUrl + "/user/new/" + userType,
      options
    );
    console.log(response);
    if (response && response.data && response.data.created === true) {
      // window.location = window.location.origin + "/profile";
      return response.data;
    }
    return response.body.error;
  } catch (error) {
    console.error(error);
    alert("Something went wrong, Please try again!");
  }
};

export default createUser;
