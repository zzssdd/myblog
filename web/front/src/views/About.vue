<template>
  <div>
    <AboutMe />
    <Contacts />
  </div>
</template>

<script>
import AboutMe from '@/components/AboutMe.vue'
import Contacts from '@/components/Contacts.vue'

export default {
  name: 'About',
  components: {
    AboutMe,
    Contacts
  },
  mounted: function () {
    this.changeTitle('关于')
  }
}
</script>
