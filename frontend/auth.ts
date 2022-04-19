import {ConditionsFailedEvent, replace} from 'svelte-spa-router'

async function checkAuthentication() {
    try {
        const res = await fetch('/api/ping',{
            method:'HEAD'
        })

        if(res.status !== 200) {
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
    checkAuthentication,
    conditionsFailed
}