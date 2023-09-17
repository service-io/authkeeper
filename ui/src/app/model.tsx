import {ReactElement} from "react";
import {
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  CardMedia,
  Checkbox,
  FormControlLabel,
  Typography
} from "@mui/material";

import onePic from "assets/one.png"
import twoPic from "assets/two.png"

const Model = (): ReactElement => {
  const cardSx = {
    width: 200,
    height: 100,
  }
  const card0Sx = {
    width: 200,
  }
  const card1Sx = {
    width: 200,
  }
  const picSx = {
    width: 200,
    height: 100,
  }
  return (
    <>
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        minHeight="100vh"
      >
        <Card sx={{width: 600, height: 600}} variant="outlined">
          <CardHeader sx={{backgroundColor: "mediumpurple", height: 20}} title="识别系统 V1.0"/>
          <CardContent>
            <Typography sx={{fontWeight: "bold"}} variant="h6" component="div" align="center">
              基于深度学习的燃气管道缺陷智能识别系统
            </Typography>
            <Box sx={{display: 'flex', alignItems: 'center', justifyContent: 'space-around'}}>
              <Box>
                <Typography align="center">
                  实时图像
                </Typography>
                <Card sx={cardSx}>
                  <CardMedia
                    component="img"
                    alt="green iguana"
                    sx={picSx}
                    image={onePic}
                  />
                </Card>
                <br/>
                <Typography align="center">
                  处理结果图像
                </Typography>
                <Card sx={cardSx}>
                  <CardMedia
                    component="img"
                    alt="green iguana"
                    sx={picSx}
                    image={twoPic}
                  />
                </Card>
              </Box>
              <Box>
                <br/>
                <Box sx={{display: 'flex', alignItems: 'center'}}>
                  <Typography sx={{writingMode: "vertical-rl"}} color="text.secondary">
                    缺陷位置
                  </Typography>
                  <Card sx={card0Sx} variant="outlined">
                    <CardContent>
                      <Typography variant="body2" gutterBottom>
                        置信度: 1
                      </Typography>
                      <Typography variant="body2" gutterBottom>
                        xmin: 0
                      </Typography>
                      <Typography variant="body2" gutterBottom>
                        ymin: 0
                      </Typography>
                      <Typography variant="body2" gutterBottom>
                        xmax: 0
                      </Typography>
                      <Typography variant="body2" gutterBottom>
                        ymax: 0
                      </Typography>
                    </CardContent>
                  </Card>
                </Box>
                <br/>
                <Box sx={{display: 'flex', alignItems: 'center'}}>
                  <Typography sx={{writingMode: "vertical-rl"}} color="text.secondary">
                    缺陷类别
                  </Typography>
                  <Card sx={card1Sx} variant="outlined">
                    <CardContent>
                      <FormControlLabel sx={{mb: -2}} control={<Checkbox size="small" />} label="裂纹缺陷" />
                      <FormControlLabel sx={{mb: -2}} disabled control={<Checkbox defaultChecked size="small" />} label="孔洞缺陷" />
                      <FormControlLabel sx={{mb: -2}} control={<Checkbox size="small" />} label="接口破损缺陷" />
                      <FormControlLabel sx={{mb: -2}} control={<Checkbox size="small" />} label="第三方破坏缺陷" />
                      <FormControlLabel control={<Checkbox size="small" />} label="其他缺陷" />
                    </CardContent>
                  </Card>
                </Box>
              </Box>
            </Box>
          </CardContent>
          <CardActions sx={{display: 'flex', alignItems: 'center', justifyContent: 'space-around'}}>
            <Button variant="outlined" size="small">开始运行</Button>
            <Button variant="outlined" size="small">暂停运行</Button>
            <Button variant="outlined" size="small">图像捕捉</Button>
            <Button variant="outlined" size="small">退出系统</Button>
            <Button variant="outlined" size="small">关于</Button>
          </CardActions>
        </Card>
      </Box>
    </>
  )
}

export default Model
