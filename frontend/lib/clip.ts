interface Clip {
    clip_id:number
    user_id:number
    creator:string
    title:string
    description:string
    views:number
    uploaded_at:Date
}

async function getClips(options:{page:number,amount:number}):Promise<Clip[]> {
    try {
        const res = await fetch(`/api/clip/get?page=${options.page}&amount=${options.amount}`)

        if(res.status !== 200) {
            console.log(res.status,res.statusText)

            return Promise.reject('Failed to get clips')
        }

        const clips = <Clip[]> (await res.json())

        return Promise.resolve(clips)
    } catch(err) {
        console.trace(err)
        
        return Promise.reject('Unknown Error')
    }
}

async function getClip(options:{clip_id:number}):Promise<Clip> {
    try {
        const res = await fetch(`/api/clip/${ options.clip_id }`)

        if(res.status !== 200) {
            console.log(res.status,res.statusText)

            return Promise.reject('Failed to get clip')
        }

        const clip = <Clip> (await res.json())

        return Promise.resolve(clip)
    } catch(err) {
        console.trace(err)

        return Promise.reject('Unknown Error')
    }
}

async function incrementViews(clipID:number):Promise<number> {
    try {
        const res = await fetch(`/api/clip/views/${ clipID }`,{method:'POST'})

        if(res.status !== 200) {
            console.log(res.status,res.statusText)

            return Promise.reject(`Failed to add to views for clip ${ clipID }`)
        }

        const views = <number> (await res.json())

        return Promise.resolve(views)
    } catch(err) {
        console.trace(err)

        return Promise.reject('Unknown Error')
    }
}

export {
    getClip,
    getClips,
    incrementViews
}
