-- ----------------------------
-- 背景图片ID
-- ----------------------------
SET @instance_background_img = 12;

-- ----------------------------
-- screen_instance（大屏实例）
-- ----------------------------
SET @version = 1;

INSERT INTO `screen_instance` (
    `add_time`,
    `add_user`,
    `edit_time`,
    `edit_user`,
    `del_flag`,
    `instance_title`,
    `instance_background_color`,
    `instance_background_img`,
    `instance_height`,
    `instance_width`,
    `instance_version`,
    `instance_view_img`
)
VALUES
(
    '2020-02-23 13:32:07',
    '1',
    '1',
    '2020-02-23 13:32:07',
    '1',
    '#263546',
    @instance_background_img,
    '1080',
    '腾讯云服务监控',
    '',
    '1920',
    @version
);

-- ----------------------------
-- \"database\":9
-- 更换为自己的ID
-- ----------------------------

SET @instance_id = @@IDENTITY;

-- ----------------------------
-- 顶部大标题
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '50',
    'chart0',
    '8000',
    '{\"title\":\"云计算服务监控\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":40,\"fontWeight\":\"normal\",\"textAlign\":\"center\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '400',
    '755',
    '40',
    @version
);
-- ----------------------------
-- 任务总量
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart1',
    '8000',
    '{\"title\":\"任务总量\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '95',
    '100',
    '125',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select x,y,legend from template_3_tta_line\",\"x\":\"x\",\"y\":\"y\",\"legend\":\"legend\",\"chartType\":\"lineStackingArea\"}',
    'lineStackingArea',
    'false',
    '[]',
    '225',
    'chart2',
    '8000',
    '{\"title\":{\"text\":\"堆叠面积图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"grid\":{\"top\":\"5%\",\"left\":\"0\",\"right\":\"5%\",\"bottom\":\"0\",\"containLabel\":true},\"tooltip\":{\"trigger\":\"axis\"},\"legend\":{\"show\":false,\"data\":[\"s1\",\"s2\"],\"orient\":\"vertical\",\"left\":0,\"top\":0,\"textStyle\":{\"color\":\"#fff\"}},\"xAxis\":{\"name\":\"\",\"type\":\"category\",\"boundaryGap\":false,\"data\":[\"01/01\",\"02/01\",\"03/01\",\"04/01\",\"05/01\",\"06/01\"],\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"},\"rotate\":0},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"yAxis\":{\"name\":\"\",\"type\":\"value\",\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitArea\":{\"show\":false,\"areaStyle\":{\"color\":[\"rgba(250,250,250,0.3)\",\"rgba(200,200,200,0.3)\"]}},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"color\":[\"#00C8DC\",\"#006EFF\"],\"series\":[{\"name\":\"s1\",\"type\":\"line\",\"data\":[1160,600,1800,400,1300,600],\"areaStyle\":{\"color\":\"rgba(0, 200, 220, 0.1)\"},\"smooth\":true},{\"name\":\"s2\",\"type\":\"line\",\"data\":[600,800,120,1200,200,500],\"areaStyle\":{\"color\":\"rgba(0, 110, 255, 0.1)\"},\"smooth\":true}],\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '415',
    '80',
    '190',
    @version
);
-- ----------------------------
-- 任务分布
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart3',
    '8000',
    '{\"title\":\"任务分布\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '95',
    '100',
    '455',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_3_td_radar\",\"value\":\"value\",\"name\":\"name\",\"max\":\"max\",\"legend\":\"legend\",\"chartType\":\"radarBasic\"}',
    'radarBasic',
    'false',
    '[]',
    '200',
    'chart4',
    '8000',
    '{\"title\":{\"text\":\"基础雷达图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"tooltip\":{},\"legend\":{\"show\":false,\"orient\":\"vertical\",\"data\":[\"s1\",\"s2\"],\"top\":0,\"left\":0,\"textStyle\":{\"color\":\"#fff\"}},\"radar\":{\"indicator\":[{\"max\":\"200\",\"name\":\"A\"},{\"max\":\"200\",\"name\":\"B\"},{\"max\":\"200\",\"name\":\"C\"},{\"max\":\"200\",\"name\":\"D\"},{\"max\":\"200\",\"name\":\"E\"},{\"max\":\"200\",\"name\":\"F\"}],\"splitNumber\":3,\"splitArea\":{\"show\":false},\"radius\":\"70\"},\"color\":[\"#FF5A32\",\"#006EFF\"],\"series\":[{\"name\":\"\",\"type\":\"radar\",\"symbol\":\"none\",\"data\":[{\"name\":\"s1\",\"value\":[\"80\",\"50\",\"100\",\"60\",\"40\",\"50\"]},{\"name\":\"s2\",\"value\":[\"150\",\"180\",\"170\",\"120\",\"140\",\"110\"]}]}],\"toolbox\":{\"feature\":{\"saveAsImage\":{}}}}',
    'false',
    '180',
    '75',
    '510',
    @version
);
-- ----------------------------
-- 任务成功
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart5',
    '8000',
    '{\"title\":\"任务成功\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '95',
    '325',
    '455',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_3_ts_pie\",\"value\":\"value\",\"chartType\":\"piePercent\"}',
    'piePercent',
    'false',
    '[]',
    '210',
    'chart6',
    '8000',
    '{\"title\":{\"text\":\"62%\",\"x\":\"center\",\"y\":\"center\",\"textStyle\":{\"color\":\"#98a0c4\",\"fontSize\":20}},\"series\":[{\"name\":\"main\",\"type\":\"pie\",\"radius\":[\"70%\",\"85%\"],\"center\":[\"50%\",\"50%\"],\"label\":{\"normal\":{\"show\":false}},\"data\":[{\"value\":62,\"itemStyle\":{\"normal\":{\"color\":\"rgba(119,96,246,1)\"}}},{\"value\":38,\"itemStyle\":{\"normal\":{\"color\":\"rgba(44,52,92,1)\"}}}]}],\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '180',
    '310',
    '510',
    @version
);
-- ----------------------------
-- 任务完成率
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart7',
    '8000',
    '{\"title\":\"任务完成率 \",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '120',
    '100',
    '755',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select x,y,legend from template_3_tcr_line\",\"x\":\"x\",\"y\":\"y\",\"legend\":\"legend\",\"chartType\":\"lineStackingArea\"}',
    'lineStackingArea',
    'false',
    '[]',
    '200',
    'chart8',
    '8000',
    '{\"title\":{\"text\":\"堆叠面积图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"grid\":{\"top\":\"5%\",\"left\":\"0\",\"right\":\"5%\",\"bottom\":\"0\",\"containLabel\":true},\"tooltip\":{\"trigger\":\"axis\"},\"legend\":{\"show\":false,\"data\":[\"s1\",\"s2\"],\"orient\":\"vertical\",\"left\":0,\"top\":0,\"textStyle\":{\"color\":\"#fff\"}},\"xAxis\":{\"name\":\"\",\"type\":\"category\",\"boundaryGap\":false,\"data\":[\"A\",\"B\",\"C\",\"D\",\"E\",\"F\",\"G\"],\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"},\"rotate\":0},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"yAxis\":{\"name\":\"\",\"type\":\"value\",\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitArea\":{\"show\":false,\"areaStyle\":{\"color\":[\"rgba(250,250,250,0.3)\",\"rgba(200,200,200,0.3)\"]}},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"series\":[{\"name\":\"s1\",\"type\":\"line\",\"symbol\":\"none\",\"data\":[1200,1100,1700,900,1200,800,700],\"itemStyle\":{\"normal\":{\"lineStyle\":{\"color\":\"rgba(0,0,0,0)\"}}},\"areaStyle\":{\"color\":\"rgba(0, 110, 255, 0.7)\"},\"smooth\":true},{\"name\":\"s2\",\"type\":\"line\",\"symbol\":\"none\",\"data\":[150,300,700,400,500,300,300],\"itemStyle\":{\"normal\":{\"lineStyle\":{\"color\":\"rgba(0,0,0,0)\"}}},\"areaStyle\":{\"color\":\"rgba(0, 200, 220, 0.7)\"},\"smooth\":true}],\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '400',
    '80',
    '815',
    @version
);
-- ----------------------------
-- 云计算任务趋势
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart9',
    '8000',
    '{\"title\":\"云计算趋势\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '120',
    '590',
    '125',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_3_cctt_bar UNION select * from template_3_cctt_line\",\"x\":\"x\",\"y\":\"y\",\"legend\":\"legend\",\"chartType\":\"histogramComplex\"}',
    'histogramComplex',
    'false',
    '{}',
    '250',
    'chart10',
    '8000',
    '{\"title\":{\"show\":false,\"text\":\"复合柱状图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"}},\"grid\":{\"top\":\"5%\",\"left\":\"0\",\"right\":\"0\",\"bottom\":\"5%\",\"containLabel\":true},\"legend\":{\"show\":false,\"data\":[\"bar\",\"line\"],\"textStyle\":{\"color\":\"#fff\"},\"top\":\"0\",\"left\":\"0\"},\"dataZoom\":{\"show\":false,\"start\":0,\"end\":100},\"xAxis\":[{\"type\":\"category\",\"boundaryGap\":true,\"data\":[\"01/01\",\"02/01\",\"03/01\",\"04/01\",\"05/01\",\"06/01\",\"07/01\",\"08/01\",\"09/01\",\"10/01\",\"11/01\",\"12/01\"],\"name\":\"\",\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"},\"rotate\":0},\"axisTick\":{\"show\":false},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\",\"type\":\"dashed\"}},\"axisPointer\":{\"type\":\"shadow\"}},{\"type\":\"category\",\"boundaryGap\":true,\"data\":[\"01/01\",\"02/01\",\"03/01\",\"04/01\",\"05/01\",\"06/01\",\"07/01\",\"08/01\",\"09/01\",\"10/01\",\"11/01\",\"12/01\"],\"name\":\"\",\"axisLabel\":{\"show\":false,\"textStyle\":{\"color\":\"#fff\"}},\"axisTick\":{\"show\":false},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisPointer\":{\"type\":\"shadow\"}}],\"yAxis\":[{\"type\":\"value\",\"scale\":false,\"name\":\"\",\"min\":0,\"boundaryGap\":[0.2,0.2],\"axisLabel\":{\"show\":false,\"textStyle\":{\"color\":\"#fff\"}},\"axisTick\":{\"show\":false},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"},\"type\":\"dashed\"}},{\"type\":\"value\",\"scale\":false,\"name\":\"\",\"min\":0,\"boundaryGap\":[0.2,0.2],\"axisLabel\":{\"show\":false,\"textStyle\":{\"color\":\"#fff\"}},\"axisTick\":{\"show\":false},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}}}],\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"color\":[\"#006EFF\",\"#6E3CFF\"],\"series\":[{\"name\":\"\",\"type\":\"bar\",\"barWidth\":\"35%\",\"xAxisIndex\":1,\"yAxisIndex\":1,\"data\":[40,20,70,100,40,20,30,45,120,60,50,15]},{\"name\":\"\",\"type\":\"line\",\"symbol\":\"none\",\"data\":[60,100,90,140,80,50,100,90,150,80,90,50]}],\"backgroundColor\":\"rgba(0,0,0,0)\"}',
    'false',
    '790',
    '570',
    '170',
    @version
);
-- ----------------------------
-- 任务管理
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart11',
    '8000',
    '{\"title\":\"任务管理\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '95',
    '590',
    '455',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_3_tm_list\",\"chartType\":\"rotationList\"}',
    'rotationList',
    'false',
    '[]',
    '530',
    'chart12',
    '8000',
    '{\"column\":{\"show\":true,\"height\":40,\"fontColor\":\"#ffffff\",\"fontSize\":15},\"columnList\":[\"时间\",\"任务1\",\"任务2\"],\"columnWidthList\":[\"150px\",\"120px\",\"120px\"],\"data\":{\"height\":40,\"fontColor\":\"#e2e2e2\",\"fontSize\":15},\"speed\":3,\"rowNum\":12}',
    'false',
    '405',
    '550',
    '500',
    @version
);


-- ----------------------------
-- 平均容量
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart13',
    '8000',
    '{\"title\":\"平均容量\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '95',
    '1015',
    '455',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_3_ac_histogram\",\"x\":\"x\",\"y\":\"y\",\"legend\":\"legend\",\"chartType\":\"histogramStacking\"}',
    'histogramStacking',
    'false',
    '{}',
    '225',
    'chart14',
    '8000',
    '{\"title\":{\"text\":\"堆叠柱状图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"grid\":{\"left\":\"0\",\"right\":\"5%\",\"bottom\":\"0\",\"containLabel\":true,\"top\":\"5%\"},\"legend\":{\"left\":\"1%\",\"top\":\"5%\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"tooltip\":{},\"dataset\":{\"source\":[[\"product\",\"data1\",\"data2\"],[\"02/02\",\"80\",\"50\"],[\"02/03\",\"50\",\"150\"],[\"02/04\",\"160\",\"120\"],[\"02/05\",\"120\",\"190\"],[\"02/06\",\"40\",\"140\"]]},\"xAxis\":{\"name\":\"\",\"type\":\"category\",\"boundaryGap\":true,\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"},\"rotate\":0},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"yAxis\":{\"name\":\"\",\"type\":\"value\",\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitArea\":{\"show\":false,\"areaStyle\":{\"color\":[\"rgba(250,250,250,0.3)\",\"rgba(200,200,200,0.3)\"]}},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"color\":[\"#006EFF\",\"#00C8DC\"],\"series\":[{\"type\":\"bar\",\"barWidth\":\"30%\"},{\"type\":\"bar\",\"barWidth\":\"30%\"}],\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '365',
    '1000',
    '500',
    @version
);

-- ----------------------------
-- 存储容量
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart15',
    '8000',
    '{\"title\":\"存储容量\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '95',
    '1015',
    '755',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_3_sc_histogram\",\"y\":\"y\",\"x\":\"x\",\"chartType\":\"histogramGradientHorizontal\"}',
    'histogramGradientHorizontal',
    'false',
    '{}',
    '250',
    'chart16',
    '8000',
    '{\"title\":{\"text\":\"水平柱状图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"tooltip\":{\"trigger\":\"axis\",\"axisPointer\":{\"type\":\"shadow\"}},\"grid\":{\"left\":\"0\",\"right\":\"5%\",\"bottom\":\"0\",\"containLabel\":true,\"top\":\"5%\"},\"legend\":{\"show\":true,\"orient\":\"vertical\",\"data\":[],\"top\":\"0\",\"left\":\"0\",\"textStyle\":{\"color\":\"#ffffff\"}},\"xAxis\":{\"data\":[],\"name\":\"\",\"type\":\"value\",\"axisLabel\":{\"show\":false,\"textStyle\":{\"color\":\"#fff\"}},\"axisTick\":{\"show\":false},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}}},\"yAxis\":{\"name\":\"\",\"type\":\"category\",\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"data\":[\"A\",\"B\",\"C\",\"D\",\"E\"]},\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"color\":[\"#006EFF\"],\"series\":{\"name\":\"\",\"type\":\"bar\",\"barWidth\":\"40%\",\"data\":[80,50,160,120,40]},\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '400',
    '1000',
    '790',
    @version
);
-- ----------------------------
-- 云计算总资源
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart17',
    '8000',
    '{\"title\":\"云计算总资源\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '150',
    '1450',
    '125',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_3_tccr_line\",\"x\":\"x\",\"y\":\"y\",\"legend\":\"legend\",\"chartType\":\"lineStacking\"}',
    'lineStacking',
    'false',
    '[]',
    '245',
    'chart18',
    '8000',
    '{\"title\":{\"text\":\"堆叠折线图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"grid\":{\"top\":\"5%\",\"left\":\"0\",\"right\":\"5%\",\"bottom\":\"0\",\"containLabel\":true},\"legend\":{\"left\":\"1%\",\"top\":\"5%\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"tooltip\":{},\"dataset\":{\"source\":[[\"product\",\"s1\",\"s2\"],[\"01/01\",\"1160\",\"600\"],[\"02/01\",\"600\",\"800\"],[\"03/01\",\"1800\",\"120\"],[\"04/01\",\"400\",\"1200\"],[\"05/01\",\"1300\",\"200\"],[\"06/01\",\"600\",\"500\"]]},\"xAxis\":{\"name\":\"\",\"type\":\"category\",\"boundaryGap\":true,\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"},\"rotate\":0},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"yAxis\":{\"name\":\"\",\"type\":\"value\",\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitArea\":{\"show\":false,\"areaStyle\":{\"color\":[\"rgba(250,250,250,0.3)\",\"rgba(200,200,200,0.3)\"]}},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"color\":[\"#00C8DC\",\"#006EFF\"],\"series\":[{\"type\":\"line\"},{\"type\":\"line\"}],\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '405',
    '1430',
    '175',
    @version
);

-- ----------------------------
-- CPU利用率
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart19',
    '8000',
    '{\"title\":\"CPU利用率\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '120',
    '1450',
    '455',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_3_cpuur_pie\",\"value\":\"value\",\"chartType\":\"piePercent\"}',
    'piePercent',
    'false',
    '[]',
    '250',
    'chart20',
    '8000',
    '{\"title\":{\"text\":\"62%\",\"x\":\"center\",\"y\":\"center\",\"textStyle\":{\"color\":\"#98a0c4\",\"fontSize\":20},\"show\":true},\"series\":[{\"name\":\"main\",\"type\":\"pie\",\"radius\":[\"60%\",\"75%\"],\"center\":[\"50%\",\"50%\"],\"label\":{\"normal\":{\"show\":false}},\"data\":[{\"value\":62,\"itemStyle\":{\"normal\":{\"color\":\"rgba(0,200,220,1)\"}}},{\"value\":38,\"itemStyle\":{\"normal\":{\"color\":\"rgba(44,52,92,1)\"}}}]}],\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '400',
    '1430',
    '500',
    @version
);

-- ----------------------------
-- 内存利用率
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart21',
    '8000',
    '{\"title\":\"内存利用率\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '120',
    '1450',
    '755',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select x,y,legend from template_3_mu_line\",\"x\":\"x\",\"y\":\"y\",\"legend\":\"legend\",\"chartType\":\"lineStackingArea\"}',
    'lineStackingArea',
    'false',
    '[]',
    '230',
    'chart22',
    '8000',
    '{\"title\":{\"text\":\"堆叠面积图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"grid\":{\"top\":\"5%\",\"left\":\"0\",\"right\":\"5%\",\"bottom\":\"0\",\"containLabel\":true},\"tooltip\":{\"trigger\":\"axis\"},\"legend\":{\"show\":false,\"data\":[\"data\"],\"orient\":\"vertical\",\"left\":0,\"top\":0,\"textStyle\":{\"color\":\"#fff\"}},\"xAxis\":{\"name\":\"\",\"type\":\"category\",\"boundaryGap\":false,\"data\":[\"A\",\"B\",\"C\",\"D\",\"E\",\"F\"],\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"},\"rotate\":0},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#627288\"}}},\"yAxis\":{\"name\":\"\",\"type\":\"value\",\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitArea\":{\"show\":false,\"areaStyle\":{\"color\":[\"rgba(250,250,250,0.3)\",\"rgba(200,200,200,0.3)\"]}},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"color\":[\"#006EFF\"],\"series\":[{\"name\":\"data\",\"type\":\"line\",\"symbol\":\"none\",\"data\":[1100,600,1800,400,1300,600],\"areaStyle\":{\"color\":\"rgba(0, 110, 255, 0.14)\"}}],\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '410',
    '1435',
    '805',
    @version
);