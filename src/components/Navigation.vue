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
        <v-alert info value="true" class="ma-3">
          좌석 한정으로 '버스 탑승'을 원하는 분은 꼭 미리 알려주세요.
        </v-alert>
        <div class="text-xs-left ma-4">
          <h4><v-icon>directions_bus</v-icon>버스 탑승 위치</h4>
          <span class="blue--text">'양재역 9번 출구'와 서초문화예술회관 사이길</span>에서 9/23(토) <span class="blue--text">'10:30분'</span>까지 버스 탑승을 합니다. 상세한 정보는 아래를 클릭해주세요.
        </div>
        <v-btn outline class="grey" @click.native.stop="handleDisplayReservation()">
          <v-icon class="icon">event_note</v-icon>
          <span class="icon-text">버스탑승 상세 위치보기 (양재역 9번 출구 방향)</span>
        </v-btn>
        <v-dialog v-model="dialogBusSeoul" lazy absolute width="90vh">
          <v-card>
            <v-card-title>
              <v-icon large>directions</v-icon><div class="headline">버스 탑승 위치</div>
            </v-card-title>
            <v-card-text>
              <img src="../assets/map.png">
              <p>양재역에서 하차하고 <span class="yellow">'9번 출구'과 '서초문화예술회관' 사이길</span>에서 버스 탑승하시면 됩니다.</p>
              <v-layout>
                <v-flex>
                  <img src="../assets/bus01.jpg">
                </v-flex>
                <v-flex>
                  <img src="../assets/bus02.jpg">
                </v-flex>
              </v-layout>
              <p>버스는 '<span class="yellow">경기78초 2019</span>, 테마고속투어(주)' 입니다.</p>
              <p>혹시 차량의 위치를 못 찾겠거나 문의가 필요할 땐<br />버스 탑승을 담당하는 '<span class="yellow">전대우(010-4116-9164)</span>' 친구에게 문자 혹은 전화주세요.</p>
              <p>확인하여 바로 연락하겠습니다.</p>
            </v-card-text>
            <v-card-title>
              <v-icon large>rotate_right</v-icon><div class="headline">서울 복귀 시간</div>
            </v-card-title>
            <v-card-text>
              <p>식사 후 <span class="yellow">4시 30분</span>까지 하차하신 위치에서 동일한 버스로 탑승하시면됩니다.</p>
              <p>서울까지 이동 시간은 약 4시간 (휴게 시간 20분 포함)을 예상하여 8시 30분 도착 예정입니다.<br />다시 한 번, 먼 곳까지 찾아와주셔서 매우 감사드립니다.</p>
            </v-card-text>
            <v-card-text>
              <v-divider class="mt-3 mb-3"></v-divider>
              <v-icon large class="mb-3">tag_face</v-icon>
              <p>하행/상행 길 오고 가실 때<br />간단한 먹거리를 준비하였으니 편안한 마음으로 함께 해주시면 감사하겠습니다.</p>
            </v-card-text>
            <v-card-actions>
              <v-btn block primary class="white--text" @click.native="dialogBusSeoul = false">닫기</v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </div>
    </v-container>
  </div>
</template>

<script>
export default {
  name: 'navigation',
  methods: {
    handleDisplayReservation() {
      this.dialogBusSeoul = true;
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
