import {ConditionsFailedEvent, replace} from 'svelte-spa-router'
import {getMe} from './lib/user'

async function isAuthenticated() {
    try {
        const user = await getMe()

        if(!user) {
            return Promise.resolve(false)
        }

        return Promise.resolve(true)
    } catch(err) {
        console.trace(err)

        return Promise.reject(err)
    }
}

function conditionsFailed(event:ConditionsFailedEvent) {
    const authRoutes = [
        '/login',
        '/signup'
    ]

    if (authRoutes.includes(<string> event.detail.route)) {
        return replace('/')
    }

    return replace('/login')
}

export {
    conditionsFailed,
    isAuthenticated
}