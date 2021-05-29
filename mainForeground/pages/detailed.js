import Head from 'next/head'
import styles from '../public/style/pages/detailed.module.css'
import Header from '../components/Header'
import servicePath from '../config/apiUrl'
import { Row, Col, Breadcrumb, Icon } from 'antd'
import axios from 'axios'


export default function Detailed(props) {
    console.log(props)
    return (
        <div>
            <Head>
                <title>Detailed</title>
                <link rel="icon" href="/favicon.ico" />
            </Head>
            <Header />
            <Row className="commmain" type="flex" justify="center">
                <Col className="commleft" xs={24} sm={24} md={16} lg={18} xl={14}>
                    <div>
                        <div className={styles.breaddiv}>
                            <Breadcrumb>
                                <Breadcrumb.Item><a href="/">首页</a></Breadcrumb.Item>
                                <Breadcrumb.Item>新闻</Breadcrumb.Item>
                            </Breadcrumb>
                        </div>
                        <div>
                            <div className={styles.ditailedtile}>
                                <h1>{props.Title}</h1>
                                <span><Icon type="calendar" /> {props.UpdatedAt.split('T')[0] + ' ' + props.UpdatedAt.split('T')[1].split('+')[0]} </span>
                                <span>{props.Author}</span>
                                <p>{props.Content}</p>

                            </div>
                        </div>
                    </div>
                </Col>
            </Row>
        </div>
    )
}

//引入后进行修改
Detailed.getInitialProps = async (context) => {

    console.log(context.query.ID)
    let id = context.query.ID
    const promise = new Promise((resolve) => {

        axios(servicePath.getArticleById + id).then(
            (res) => {
                console.log(res.data)
                resolve(res.data.data)
            }
        )
    })

    return await promise
}