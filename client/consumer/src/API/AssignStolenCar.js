import fetchData from "./../Utils/FetchData";
import Constants from "./../Constants/Constants";
import Url from "./../Config/Config";

const assignUnassignedStolenCar = async uID => {
  const options = {
    method: Constants.httpMethods.POST,
    headers: {
      "Content-Type": "application/json"
    }
  };
  try {
    const response = await fetchData(
      Url.apiUrl + "/stolen/car/assign/" + uID,
      options
    );
    console.log(response);
    if (response && response.data && response.data.created === true) {
      // window.location = window.location.origin + "/profile";
      return response.data;
    } else if (response && response.data && response.data.created === false) {
      return response;
    }
    return response.body.error;
  } catch (error) {
    console.error(error);
    alert("Something went wrong, Please try again!");
  }
};

export default assignUnassignedStolenCar;
