import fetchData from "./../Utils/FetchData";
import Constants from "./../Constants/Constants";
import Url from "./../Config/Config";

const changeStolenCarStatus = async requestObj => {
  const options = {
    method: Constants.httpMethods.PATCH,
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(requestObj)
  };
  try {
    const response = await fetchData(
      Url.apiUrl + "/stolen/car/status",
      options
    );
    console.log(response);
    if (response && response.data && response.data.created === true) {
      return response.data;
    }
    return response.body.error;
  } catch (error) {
    console.error(error);
    alert("Something went wrong, Please try again!");
  }
};

export default changeStolenCarStatus;
