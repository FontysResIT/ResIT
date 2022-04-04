import axios from 'axios'

    const BASE_URL = 'https://localhost:8080/api'

    const headers = {
            // here we put our headers
    };
class AxiosService {
  getAllReservations(){
    try {
        await axios.get(`${BASE_URL}/reservations`, {
          headers: headers,
      })
      .then(function (response) {
        return response
      })
    } catch (error) {
        console.log(error)
    }
  }
  getReservationById(id){
    try {
        await axios.get(`${BASE_URL}/reservations/${id}`, {
          headers: headers,
          params: {id: id}
      })
      .then(function (response) {
        return response
      })
    } catch (error) {
        console.log(error)
    }
  }
  getReservationsByDate(date){
    try {
        await axios.get(`${BASE_URL}/reservations/date/${date}`, {
          headers: headers
      })
      .then(function (response) {
        return response
      })
    } catch (error) {
        console.log(error)
    }
  }
  createReservation(reservation){
   try {
     await axios.post(`${BASE_URL}/reservations`, { reservation: reservation}, {
      headers: headers
    })
    .then(function (response) {
      return response
    })
   } catch (error) {
     console.log(error)
   }
  }
  cancelReservationById(id){
   try {
     const reservation = this.getReservationById(id)
     reservation.cancelled = true

     await axios.put(`${BASE_URL}/reservations`, { reservation: reservation}, {
     headers: headers
   })
   .then(function (response) {
     return response
   })

   } catch (error) {
     console.log(error)
   }
  }
  getAllTimeSlots(){
    try {
      await axios.get(`${BASE_URL}/timeSlots`, {
        headers: headers,
    })
    .then(function (response) {
      return response
    })
  } catch (error) {
      console.log(error)
  }
  }
  getTimeSlotsByDate(date){
    try {
      await axios.get(`${BASE_URL}/timeSlots/date/${date}`, {
        headers: headers
    })
    .then(function (response) {
      return response
    })
  } catch (error) {
      console.log(error)
  }
  }
  getOccopiedTimeSlotsByDate(date){
    try {
      await axios.get(`${BASE_URL}/timeSlots/date/${date}/occupied`, {
        headers: headers
    })
    .then(function (response) {
      return response
    })
  } catch (error) {
      console.log(error)
  }
  }
  createTimeSlot(timeSlot){
    try {
      await axios.post(`${BASE_URL}/timeSlots`, { timeSlot: timeSlot}, {
       headers: headers
     })
     .then(function (response) {
       return response
     })
    } catch (error) {
      console.log(error)
    }
  }
  getApplicationConfiguration(){
    try {
      await axios.get(`${BASE_URL}/appconfig`, {
        headers: headers,
    })
    .then(function (response) {
      return response
    })
  } catch (error) {
      console.log(error)
  }
  }
  saveApplicationConfiguration(appconfig){ 
    try{
      await axios.put(`${BASE_URL}/appconfig`, { appconfig: appconfig}, {
      headers: headers
      })
      .then(function (response) {
      return response
      })
    } catch (error) {
      console.log(error)
    }
  }
  ToggleRemindPreferences(id){
    try {
      const reservation = this.getReservationById(id)
      reservation.remind_preferences = !reservation.remind_preferences
 
      await axios.put(`${BASE_URL}/reservations`, { reservation: reservation}, {
       headers: headers
    })
    .then(function (response) {
      return response
    })
    } catch (error) {
      console.log(error)
    }
  }
}
export default new AxiosService();