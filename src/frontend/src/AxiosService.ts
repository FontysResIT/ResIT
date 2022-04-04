import axios from "axios";
import type { Reservation } from "./models/Reservation";

const BASE_URL = "https://localhost:8080/api";

const headers = {
  // here we put our headers
};
class AxiosService {
  async getAllReservations() {
    try {
      await axios
        .get(`${BASE_URL}/reservations`, {
          headers: headers,
        })
        .then(function (response) {
          return response;
        });
    } catch (error) {
      console.log(error);
    }
  }
  async getReservationById(id: string) {
    try {
      return await axios
        .get(`${BASE_URL}/reservations/${id}`, {
          headers: headers,
          params: { id: id },
        })
        .then(function (response) {
          return response;
        });
    } catch (error) {
      console.log(error);
    }
  }
  async getReservationsByDate(date: Date) {
    try {
      await axios
        .get(`${BASE_URL}/reservations/date/${date}`, {
          headers: headers,
        })
        .then(function (response) {
          return response;
        });
    } catch (error) {
      console.log(error);
    }
  }
  async createReservation(reservation: Reservation) {
    try {
      await axios
        .post(
          `${BASE_URL}/reservations`,
          { reservation: reservation },
          {
            headers: headers,
          }
        )
        .then(function (response) {
          return response;
        });
    } catch (error) {
      console.log(error);
    }
  }
  async cancelReservationById(id: string) {
    try {
      const reservation = await this.getReservationById(id);
      (reservation as unknown as Reservation).isCancelled = true;

      await axios
        .put(
          `${BASE_URL}/reservations`,
          { reservation: reservation },
          {
            headers: headers,
          }
        )
        .then(function (response) {
          return response;
        });
    } catch (error) {
      console.log(error);
    }
  }
  async getAllTimeSlots() {
    try {
      await axios
        .get(`${BASE_URL}/timeSlots`, {
          headers: headers,
        })
        .then(function (response) {
          return response;
        });
    } catch (error) {
      console.log(error);
    }
  }
  async getTimeSlotsByDate(date: Date) {
    try {
      await axios
        .get(`${BASE_URL}/timeSlots/date/${date}`, {
          headers: headers,
        })
        .then(function (response) {
          return response;
        });
    } catch (error) {
      console.log(error);
    }
  }
  async getOccopiedTimeSlotsByDate(date: Date) {
    try {
      await axios
        .get(`${BASE_URL}/timeSlots/date/${date}/occupied`, {
          headers: headers,
        })
        .then(function (response) {
          return response;
        });
    } catch (error) {
      console.log(error);
    }
  }
  async createTimeSlot(timeSlot: TimeSlot) {
    try {
      await axios
        .post(
          `${BASE_URL}/timeSlots`,
          { timeSlot: timeSlot },
          {
            headers: headers,
          }
        )
        .then(function (response) {
          return response;
        });
    } catch (error) {
      console.log(error);
    }
  }
  async getApplicationConfiguration() {
    try {
      await axios
        .get(`${BASE_URL}/appconfig`, {
          headers: headers,
        })
        .then(function (response) {
          return response;
        });
    } catch (error) {
      console.log(error);
    }
  }
  async saveApplicationConfiguration(appconfig: { [key: string]: unknown }) {
    try {
      await axios
        .put(
          `${BASE_URL}/appconfig`,
          { appconfig: appconfig },
          {
            headers: headers,
          }
        )
        .then(function (response) {
          return response;
        });
    } catch (error) {
      console.log(error);
    }
  }
  async ToggleRemindPreferences(id: string) {
    try {
      const reservation = await this.getReservationById(id);
      // eslint-disable-next-line @typescript-eslint/ban-ts-comment
      //@ts-ignore
      reservation.remind_preferences = !reservation.remind_preferences;

      await axios
        .put(
          `${BASE_URL}/reservations`,
          { reservation: reservation },
          {
            headers: headers,
          }
        )
        .then(function (response) {
          return response;
        });
    } catch (error) {
      console.log(error);
    }
  }
}
export default new AxiosService();
