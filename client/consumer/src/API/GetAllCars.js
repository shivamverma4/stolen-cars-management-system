import fetchData from "./../Utils/FetchData";
import Constants from "./../Constants/Constants";
import Url from "./../Config/Config";

const getAllStolenCars = async (uID, userType) => {
  const options = {
    method: Constants.httpMethods.GET,
    headers: {
      "Content-Type": "application/json"
    }
  };
  try {
    const response = await fetchData(
      Url.apiUrl + "/stolen/cars/" + uID + "/" + userType,
      options
    );

    if (response && response.data) {
      return response;
    }
    return response.body;
  } catch (error) {
    console.error(error);
    alert("Something went wrong, Please try again!");
  }
};

export default getAllStolenCars;
