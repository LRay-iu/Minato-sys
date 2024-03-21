<template>
    <!-- 左半边 -->
    <el-container class="Page" style="width:900px;" direction="vertical">
        <!-----轮播图-->
        <div style="width: 888px;">
            <el-carousel height="auto" autoplay>
                <el-carousel-item v-for="item in img" :key="item" style="height: 450px">
                    <el-image style="width: 888px; height: 450px" :src="item.URL" />
                </el-carousel-item>
            </el-carousel>
        </div>
        <!--推送-->


    </el-container>
    <!-- 右半边 -->
    <el-container class="Page" style="width:450px; " direction="vertical">
        1515
    </el-container>

</template>

<script lang='ts' setup name='HomePage'>
const img = [
    { URL: "/src/assets/Carousel/10001.png" },
    { URL: "/src/assets/Carousel/10002.png" },
    { URL: "/src/assets/Carousel/10003.png" },
    { URL: "/src/assets/Carousel/10004.png" },
    { URL: "/src/assets/Carousel/10005.png" },
]
</script>

<style scoped>
.Page {
    height: 100vh;
    margin: 0 6px;
    border: 1px solid red;
    align-items: center;
}
</style>