<template>
  <v-card flat>
    <v-card-media :src="gallerySrc" height="50vh">
    </v-card-media>
    <v-card-title primary-title>
      <div>
        <h3 class="headline mb-0">{{galleryTitle}}</h3>
        <div>{{galleryDesc}}</div>
      </div>
    </v-card-title>
    <v-card-actions>
      <div class="text-xs-center">
        <v-btn outline v-on:click="prev()">
          <v-icon light left>fa-angle-left</v-icon>
          PREV
        </v-btn>
        <v-btn outline v-on:click="next()">
          NEXT
          <v-icon light right>fa-angle-right</v-icon>
        </v-btn>
      </div>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: 'gallery',
  methods: {
    next() {
      this.photoIndex = this.photoIndex + 1;
      if (this.photoIndex >= this.photos.length) {
        this.photoIndex = 0;
      }
      this.refreshGallery();
    },
    prev() {
      this.photoIndex = this.photoIndex - 1;
      if (this.photoIndex < 0) {
        this.photoIndex = this.photos.length - 1;
      }
      this.refreshGallery();
    },
    refreshGallery() {
      const photo = this.photos[this.photoIndex];
      this.gallerySrc = `/static/photos/${photo.src}`;
      this.galleryTitle = photo.title;
      this.galleryDesc = photo.desc;
    },
  },
  created() {
    this.refreshGallery();
  },
  data() {
    return {
      gallerySrc: '',
      galleryTitle: 'Hello',
      galleryDesc: 'Hello',
      photoIndex: 0,
      photos: [
        { src: '01.jpg', title: '안녕하세요.', desc: '방동근v김애진 결혼합니다. 저희의 결혼 축하에 너~무 감사드립니다.' },
        { src: '02.jpg', title: '어색어색~', desc: '오랜만에 연락드리는 분들께 어색함을 표현한 사진입니다.' },
        { src: '03.jpg', title: '촌스러운, 아니 클래식한', desc: '클래식한 매력이 있는 모습이에요.' },
        { src: '04.jpg', title: '훗훗', desc: '즐거운, 행복한 표정으로' },
        { src: '05.jpg', title: '앗!', desc: '처음보는 사람이라고 생각하실 수도 있겠지만' },
        { src: '06.jpg', title: '맞습니다. 맞고요', desc: '방동근v김애진 맞습니다!' },
        { src: '07.jpg', title: '행복하게', desc: '살아보겠습니다. 사진으로 기억도 많이 남기고요.' },
        { src: '08.jpg', title: '부드럽게', desc: '즐겁게 행복하게' },
        { src: '09.jpg', title: '그렇게, 계속', desc: '살아가 보겠습니다!' },
      ],
    };
  },
};
</script>
