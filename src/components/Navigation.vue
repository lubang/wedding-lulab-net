<template>
  <div class="location">
    <img src="../assets/navigation.png"></img>
    <v-container>
      <div class="location-title">씨엘드포레 (T. 053-767-2330)</div>
      <div class="location-detail">대구광역시 달성군 가창면 가창로 891</div>
      <div class="text-xs-center navigation-buttons">
        <v-btn dark class="green" href="http://naver.me/59zi5GRA">
          <v-icon class="icon">directions_car</v-icon>
          <span class="icon-text">네이버 길찾기</span>
        </v-btn>
        <v-btn class="yellow" href="http://dmaps.kr/6zu7v">
          <v-icon class="icon">place</v-icon>
          <span class="icon-text">다음 길찾기</span>
        </v-btn>
        <div class="bus">
          ---
        </div>
        <v-btn outline class="grey" @click.native.stop="handleDisplayReservation()">
          <v-icon class="icon">event_note</v-icon>
          <span class="icon-text">버스 탑승 신청 (서울)</span>
        </v-btn>
        <v-dialog v-model="dialogBusSeoul" lazy absolute>
          <v-card>
            <v-card-title>
              <div class="headline">버스 탑승 신청 (서울)</div>
            </v-card-title>
            <v-card-text>
              <v-form ref="reservation">
                <v-text-field
                  label="성함 (가명 적으셔도 됩니다^^)"
                  v-model="name"
                  :rules="nameRules"
                  required
                ></v-text-field>
                <v-text-field
                  label="전화번호 (버스 안내 용도)"
                  v-model="mobile"
                  :rules="mobileRules"
                  required
                ></v-text-field>
              </v-form>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn class="green--text darken-1" @click.native="dialogBusSeoul = false">닫기</v-btn>
              <v-btn class="green--text darken-1" @click="bookBusFromSeoul()">확인</v-btn>
            </v-card-actions>
            <v-card-text v-if="isSuccess">
              <v-alert success value="true">
                성공적으로 등록하였습니다. 결혼식 날 뵈요^^!
              </v-alert>
            </v-card-text>
          </v-card>
        </v-dialog>
        <v-alert info value="true" class="ma-3">
          '버스 탑승'을 신청하시면 상세 탑승 위치를 '9/20(수)'에 문자로 전달하겠습니다.
        </v-alert>
        <div class="text-xs-left ma-4">
          <h4><v-icon>directions_bus</v-icon>버스 탑승 위치</h4>
          <span class="blue--text">9/23(토) 10:30분 강남역</span> 부근에서 출발할 예정입니다. 버스 기사님과 정확한 장소가 정해지면 문자로 다시 전달하겠습니다.
        </div>
      </div>
    </v-container>
  </div>
</template>

<script>
import reservationApi from '../api/reservation';

export default {
  name: 'navigation',
  methods: {
    handleDisplayReservation() {
      this.isSuccess = false;
      this.dialogBusSeoul = true;
    },
    bookBusFromSeoul() {
      if (this.$refs.reservation.validate()) {
        const payload = {
          name: this.name,
          mobile: this.mobile,
        };
        reservationApi.save(payload)
          .then(() => {
            this.isSuccess = true;
          });
      }
    },
  },
  data() {
    return {
      dialogBusSeoul: false,
      dialogBusDaejeon: false,
      name: '',
      nameRules: [
        v => !!v || '이름은 꼭 적어주셔야 합니다.',
      ],
      mobile: '',
      mobileRules: [
        v => !!v || '전화번호는 꼭 적어주셔야 합니다.',
      ],
      isSuccess: false,
    };
  },
};
</script>

<style scoped>
.location {
  background-color: white;
  padding-top: 50px;
  padding-bottom: 50px;
  font-size: 1.1em;
}
.location-title {
  font-size: 1.30em;
  padding: 2px 5px 5px 8px;
  color: #5C5C5C;
}
.location-detail {
  font-size: 1.15em;
  padding: 2px 5px 5px 8px;
  color: #5C5C5C;
}
.navigation-buttons {
  margin-top: 10px;
}
.icon {
  font-size: 1.0em;
  margin-right: 5px;
}
.icon-text {
  font-weight: bold;
}
.bus {
  margin-top: 40px;
  margin-bottom: 10px;
}
</style>
