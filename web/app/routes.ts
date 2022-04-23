import {wrap} from 'svelte-spa-router/wrap'
import {isAuthenticated} from './auth'

import Home from './svelte/Home.svelte'
import Login from './svelte/Login.svelte'
import Signup from './svelte/Signup.svelte'
import Upload from './svelte/Upload.svelte'
import Clip from './svelte/Clip.svelte'

export default {
    '/':wrap({
        component:Home,
        conditions:[
            async () => {
                return await isAuthenticated()
            }
        ]
    }),
    '/login':wrap({
        component:Login,
        conditions:[
            async () => {
                return !(await isAuthenticated())
            }
        ]
    }),
    '/signup':wrap({
        component:Signup,
        conditions:[
            async () => {
                return !(await isAuthenticated())
            }
        ]
    }),
    '/upload':wrap({
        component:Upload,
        conditions:[
            async () => {
                return await isAuthenticated()
            }
        ]
    }),
    '/clip/:clip_id':wrap({
        component:Clip,
        conditions:[
            async () => {
                return await isAuthenticated()
            }
        ]
    })
}