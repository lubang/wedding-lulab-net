<template>
  <div class="location">
    <img src="../assets/navigation.png"></img>
    <v-container>
      <div class="location-title">씨엘드포레 (T. 053-767-2330)</div>
      <div class="location-detail">대구광역시 달성군 가창면 가창로 891</div>
      <div class="text-xs-center navigation-buttons">
        <v-btn dark class="green" href="http://naver.me/59zi5GRA">
          <v-icon class="icon">fa-map-marker</v-icon>
          <span class="icon-text">네이버 길찾기</span>
        </v-btn>
        <v-btn class="yellow" href="http://dmaps.kr/6zu7v">
          <v-icon class="icon">fa-map-pin</v-icon>
          <span class="icon-text">다음 길찾기</span>
        </v-btn>
        <div class="bus">
          <v-icon large>fa-bus</v-icon>
        </div>
        <v-btn outline class="grey" @click.native.stop="dialogBusSeoul = true">
          <v-icon class="icon">fa-ticket</v-icon>
          <span class="icon-text">버스 (서울 출발)</span>
        </v-btn>
        <v-dialog v-model="dialogBusSeoul" lazy absolute>
          <v-card>
            <v-card-title>
              <div class="headline">버스 (서울 출발) 예약</div>
            </v-card-title>
            <v-card-text>본 예약은 9/8일 이후 활성화됩니다.<br />감사합니다.</v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn class="green--text darken-1" flat="flat" @click.native="dialogBusSeoul = false">닫기</v-btn>
              <v-btn class="green--text darken-1" flat="flat" @click.native="dialogBusSeoul = false">확인</v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
        <v-alert info value="true" class="ma-3">
          <p>감사합니다.</p><p>버스 탑승은 9/8(금) 이후<br />활성화 됩니다.</p>
        </v-alert>
      </div>
    </v-container>
  </div>
</template>

<script>
export default {
  name: 'navigation',
  methods: {
    bookBusFromSeoul() {
    },
  },
  data() {
    return {
      dialogBusSeoul: false,
      dialogBusDaejeon: false,
      test: true,
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
