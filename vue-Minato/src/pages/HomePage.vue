<template>
    <PageTop />
    <el-container class="Maincontent" direction="vertical">
        <el-container direction="horizontal">
            <!-- 左半边 -->
            <el-container class="Page" style="width:1050px;" direction="vertical">
                <!-----轮播图-->
                <div style="width: 900px;">
                    <el-carousel height="auto" autoplay>
                        <el-carousel-item v-for="item in img" :key="item" style="height: 450px">
                            <el-image style="width: 900px; height: 480px" :src="item.URL" />
                        </el-carousel-item>
                    </el-carousel>
                </div>
                <el-divider class="divide" />
                <!--推送-->
                <el-container class="articleList" direction="horizontal">
                    <el-container>
                        <el-image class="thumb" src="/src/assets/article/thumb/test.png" />
                    </el-container>
                    <el-container direction="vertical" style="padding: 10px;">
                        <el-container class="title">
                            关于7旬老年人飙车硬闯巴音布鲁克导致摔下悬崖
                        </el-container>
                        <el-divider />
                        <el-container class="intro">
                            一名七旬老年人在驾车时失控，冲破路障，勇闯巴音布鲁克风景区，不幸车辆失控摔下悬崖。事发时，老人疑似因疾病影响而无法控制车辆，导致惨剧发生。事故引起了社会广泛关注，也提醒了公众对于驾驶行为的重视和对老年驾驶者的关爱。在处理此类事件时，应加强对老年人的交通安全教育和心理支持...
                        </el-container>
                    </el-container>
                </el-container>
                <!-- 复制几份凑合着用用 -->
                <el-container class="articleList" direction="horizontal">
                    <el-container>
                        <el-image class="thumb" src="/src/assets/article/thumb/test.png" />
                    </el-container>
                    <el-container direction="vertical" style="padding: 10px;">
                        <el-container class="title">
                            关于7旬老年人飙车硬闯巴音布鲁克导致摔下悬崖
                        </el-container>
                        <el-divider />
                        <el-container class="intro">
                            一名七旬老年人在驾车时失控，冲破路障，勇闯巴音布鲁克风景区，不幸车辆失控摔下悬崖。事发时，老人疑似因疾病影响而无法控制车辆，导致惨剧发生。事故引起了社会广泛关注，也提醒了公众对于驾驶行为的重视和对老年驾驶者的关爱。在处理此类事件时，应加强对老年人的交通安全教育和心理支持...
                        </el-container>
                    </el-container>
                </el-container>
                <el-container class="articleList" direction="horizontal">
                    <el-container>
                        <el-image class="thumb" src="/src/assets/article/thumb/test.png" />
                    </el-container>
                    <el-container direction="vertical" style="padding: 10px;">
                        <el-container class="title">
                            关于7旬老年人飙车硬闯巴音布鲁克导致摔下悬崖
                        </el-container>
                        <el-divider />
                        <el-container class="intro">
                            一名七旬老年人在驾车时失控，冲破路障，勇闯巴音布鲁克风景区，不幸车辆失控摔下悬崖。事发时，老人疑似因疾病影响而无法控制车辆，导致惨剧发生。事故引起了社会广泛关注，也提醒了公众对于驾驶行为的重视和对老年驾驶者的关爱。在处理此类事件时，应加强对老年人的交通安全教育和心理支持...
                        </el-container>
                    </el-container>
                </el-container>
                <el-container class="articleList" direction="horizontal">
                    <el-container>
                        <el-image class="thumb" src="/src/assets/article/thumb/test.png" />
                    </el-container>
                    <el-container direction="vertical" style="padding: 10px;">
                        <el-container class="title">
                            关于7旬老年人飙车硬闯巴音布鲁克导致摔下悬崖
                        </el-container>
                        <el-divider />
                        <el-container class="intro">
                            一名七旬老年人在驾车时失控，冲破路障，勇闯巴音布鲁克风景区，不幸车辆失控摔下悬崖。事发时，老人疑似因疾病影响而无法控制车辆，导致惨剧发生。事故引起了社会广泛关注，也提醒了公众对于驾驶行为的重视和对老年驾驶者的关爱。在处理此类事件时，应加强对老年人的交通安全教育和心理支持...
                        </el-container>
                    </el-container>
                </el-container>
                <el-container class="articleList" direction="horizontal">
                    <el-container>
                        <el-image class="thumb" src="/src/assets/article/thumb/test.png" />
                    </el-container>
                    <el-container direction="vertical" style="padding: 10px;">
                        <el-container class="title">
                            关于7旬老年人飙车硬闯巴音布鲁克导致摔下悬崖
                        </el-container>
                        <el-divider />
                        <el-container class="intro">
                            一名七旬老年人在驾车时失控，冲破路障，勇闯巴音布鲁克风景区，不幸车辆失控摔下悬崖。事发时，老人疑似因疾病影响而无法控制车辆，导致惨剧发生。事故引起了社会广泛关注，也提醒了公众对于驾驶行为的重视和对老年驾驶者的关爱。在处理此类事件时，应加强对老年人的交通安全教育和心理支持...
                        </el-container>
                    </el-container>
                </el-container>

            </el-container>
            <!-- 右半边 -->
            <el-container style="width:400px; " direction="vertical">
                <div class="recommend" style="background-color: #ffffea;">
                    <div class="top"> 今日推荐</div>
                    <div v-for="(content, index) in recommend" :key="index" class="link">
                        <el-link href="/" :underline="false" style="font-size: 18px;">{{ index + 1 }}&nbsp;.&nbsp;{{
                            content
                        }}</el-link>
                    </div>
                </div>
            </el-container>

        </el-container>
    </el-container>
    <PageBotton />
</template>

<script lang='ts' setup name='HomePage'>
import PageTop from "@/components/PageTop.vue"
import PageBotton from "@/components/PageBotton.vue"


const img = [
    { URL: "/src/assets/Carousel/10001.png" },
    { URL: "/src/assets/Carousel/10002.png" },
    { URL: "/src/assets/Carousel/10003.png" },
    { URL: "/src/assets/Carousel/10004.png" },
    { URL: "/src/assets/Carousel/10005.png" },
]

const recommend = [
    "特斯拉宣布推出新一代电动汽车",
    "汽车制造商面临供应链瓶颈挑战",
    "全球汽车销量预计将创下新纪录",
    "自动驾驶技术取得重大突破",
    "电动汽车充电基础设施建设迎来新进展",
    "某汽车公司因安全隐患召回数百万辆车",
    "车联网技术助力提升驾驶体验",
    "新型轻量化材料改善汽车燃油效率",
    "汽车行业迎来混合动力车型的兴起",
]
</script>

<style scoped>
/* 正文 */
.Maincontent {
    /* border: 1px solid red; */
    justify-content: center;
    align-items: center;
    margin-top: 10px;
    /* height: 100vh; */
}

.Page {
    /* height: 100vh; */
    margin: 0 6px;
    margin-top: 20px;
    /* border: 1px solid red; */
    align-items: center;
}

.divide {
    width: 85%;
    background: 0 0;
    /* border-top: 1px dashed #e8eaec; */
}

.articleList {
    border-radius: 10px;
    margin: 5px;
    height: 200px;
    border: solid 1px #cccccc;
    justify-content: center;
    align-items: center;
}

.articleList .thumb {
    /* align-items: center; */
    margin: 10px;
    border-radius: 10px;
    width: 188px;
    height: 120px;
    object-fit: fill;
}

.articleList .title {
    width: 666px;
    height: 30px;
    font-size: larger;
    font-weight: 600;
}

.articleList .intro {
    width: 666px;
    height: 118px;
    color: rgba(0, 0, 0, 0.5)
}

.articleList .el-divider--horizontal {
    margin: 10px 0;
    width: 100%;
    background: 0 0;
    /* border-top: 1px dashed #e8eaec; */
}

.recommend {
    border-radius: 12px;
    margin-top: 30px;
    height: 450px;
    border: solid 1px rgba(255, 153, 0, 0.38);
}

.recommend .top {
    border-radius: 10px 10px 0 0;
    margin-top: 0px;
    height: 40px;
    background-color: #FF9900;
    /* border: solid 1px #FF9900; */
    color: white;
    display: flex;
    justify-content: center;
    align-items: center;
}

.recommend .link {
    color: rgb(109, 120, 118);
    margin: 13px;
}

.el-link:hover {
    color: #ff7300;
    /* 悬停时链接的颜色 */
}
</style>