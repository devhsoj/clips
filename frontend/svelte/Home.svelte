<!-- svelte-ignore a11y-media-has-caption -->

<script lang="ts">
    import { getClips, incrementViews } from '../lib/clip'
    import { link } from 'svelte-spa-router'

    let container = null

    let clips = []
    let page = 0
    let amount = 10
    let alreadyLoadingClips = false

    async function handleScroll() {
        try {
            if(container.scrollTop + container.clientHeight === container.scrollHeight && !alreadyLoadingClips) {

                alreadyLoadingClips = true

                const newClips = await getClips({
                    page:page + 1,
                    amount
                })

                alreadyLoadingClips = false

                if(newClips === null) {
                    return
                }

                clips = clips.concat(newClips)
                page = page + 1
            }
        } catch (err) {
            console.trace(err)
        }
    }

    async function loadClips() {
        try {
            clips = await getClips({page,amount})
        } catch(err) {
            console.trace(err)
        }
    }

    async function onClipEnd(clip_id) {
        try {
            await incrementViews(clip_id)

            for(let i = 0; i < clips.length; i++) {

                if (clip_id === clips[i].clip_id) {
                    clips[i].views += 1

                    return
                }

            }
        } catch(err) {
            console.trace(err)
        }
    }
</script>

<div class="content-wrapper" on:load={loadClips()}>
    <section class="wrapper bg-soft-primary">
        <div class="container pt-5 pb-5 text-center">
            <div class="row">
                <div class="col-md-9 col-lg-7 col-xl-6 mx-auto">
                    <h2 class="display-1 mb-3">Project Clips <span class="small">by Joshua Benfield</span></h2>
<!--                    <p class="lead px-xxl-10">Check out some of the clips users have uploaded!</p>-->
                </div>
            </div>
        </div>
    </section>

    <div class="container">
        <div class="row gx-md-11 gx-lg-0">
            <div id="clipsContainer" class="col-xl-9 order-xl-2">
                <section id="clips" class="wrapper pt-8 pb-2" bind:this={container} on:scroll={handleScroll}>
                    <h2 class="display-5">Latest Clips</h2>
                    <p>Scroll down to see more clips!</p>
                    <hr class="my-4" />

                    {#if clips?.length === 0}
                        <p>Loading latest clips...</p>
                    {:else if clips === null}
                        <p>No clips found.</p>
                    {:else}
                        {#each clips as clip}
                            <section class="clip">
                                <h3 class="clip-link"><a href="/clip/{clip.clip_id}" use:link style="color:inherit;text-decoration:inherit;">{clip.title}</a></h3>
                                <p class="small" style="color:grey;">
                                    {clip.description} - <strong>{clip.creator}</strong>
                                </p>
                                <p class="small">
                                    {clip.views} Views
                                </p>
                                <hr class="my-4" />

                                <video class="player" on:ended={onClipEnd(clip.clip_id)} preload="metadata" controls>
                                    <source type={clip.type} src="/api/clip/view/{clip.clip_id}">
                                </video>
                            </section>
                            <hr class="my-8" />
                        {/each}
                    {/if}
                </section>
            </div>
            <aside class="col-xl-3 sidebar sticky-sidebar mt-md-0 py-16 d-none d-xl-block">
                <nav>
                    <ul class="nav list-unstyled lh-lg text-reset d-block">
                        <li class="nav-item"><a class="nav-link" href="/" use:link>Home</a></li>
                        <li class="nav-item"><a class="nav-link" href="/upload" use:link>Upload</a></li>
                    </ul>
                </nav>
            </aside>
        </div>
    </div>
</div>

<style>
    section::-webkit-scrollbar {
        display:none;
    }

    #clips {
        width:800px;
        height:700px;
        overflow-y:scroll;
        scrollbar-width:none;
    }

    #clips::-webkit-scrollbar {
        display:none;
    }

    .clip-link:hover {
        color:#3f78e0;
        cursor:pointer;
    }
</style>