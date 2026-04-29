---
name: Simple Blog Admin
description: A monochrome publishing workspace for a personal blog.
colors:
  ink: "#161719"
  ink-strong: "#111214"
  graphite: "#222325"
  slate: "#5F6368"
  mist: "#80858B"
  canvas: "#F2F2F0"
  canvas-elevated: "#F7F7F5"
  surface: "#FBFBF9"
  surface-muted: "#EBEBE8"
  border-subtle: "#1112141A"
  border-strong: "#1112142E"
typography:
  display:
    fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", sans-serif'
    fontSize: "42px"
    fontWeight: 600
    lineHeight: 1.05
    letterSpacing: "normal"
  headline:
    fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", sans-serif'
    fontSize: "31px"
    fontWeight: 600
    lineHeight: 1.1
    letterSpacing: "normal"
  title:
    fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", sans-serif'
    fontSize: "20px"
    fontWeight: 600
    lineHeight: 1.2
    letterSpacing: "normal"
  body:
    fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", sans-serif'
    fontSize: "14px"
    fontWeight: 400
    lineHeight: 1.6
    letterSpacing: "normal"
  label:
    fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", sans-serif'
    fontSize: "11px"
    fontWeight: 600
    lineHeight: 1.2
    letterSpacing: "0.08em"
rounded:
  sm: "8px"
  md: "10px"
  lg: "14px"
  xl: "18px"
spacing:
  xs: "8px"
  sm: "12px"
  md: "16px"
  lg: "20px"
  xl: "24px"
  2xl: "28px"
  3xl: "40px"
components:
  button-primary:
    backgroundColor: "{colors.graphite}"
    textColor: "{colors.surface}"
    rounded: "{rounded.md}"
    height: "34px"
    padding: "0 14px"
  button-secondary:
    backgroundColor: "{colors.surface}"
    textColor: "{colors.graphite}"
    rounded: "{rounded.md}"
    height: "34px"
    padding: "0 14px"
  surface-card:
    backgroundColor: "{colors.surface}"
    textColor: "{colors.ink}"
    rounded: "{rounded.xl}"
    padding: "{spacing.xl}"
  input-default:
    backgroundColor: "{colors.surface}"
    textColor: "{colors.ink}"
    rounded: "{rounded.md}"
    height: "40px"
    padding: "0 12px"
  nav-item-active:
    backgroundColor: "{colors.graphite}"
    textColor: "{colors.surface}"
    rounded: "{rounded.lg}"
    padding: "10px 12px"
---

## Overview

**Creative North Star: "The Monochrome Writing Desk"**

This admin interface should feel like a quiet editorial desk, not a colorful SaaS dashboard. The visual system is intentionally monochrome, using softened paper whites and ink-toned neutrals so the interface recedes and the writing work stays central.

Density is moderate, not sparse and not data-heavy. Surfaces are layered with light borders and restrained shadows rather than saturated fills. The overall mood is calm, practical, and slightly print-minded: closer to a polished writing room than an operations console.

The system explicitly rejects loud brand accents, oversized radii, glossy gradients, and decorative dark-mode theatrics. It should never feel neon, playful, or over-produced.

**Key Characteristics:**
- Monochrome palette with softened black and softened white, never harsh pure contrast.
- Medium-density product layout with clear sectioning and stable rhythm.
- Rounded corners are present but restrained, enough to soften the UI without making it toy-like.
- Elevation is ambient, not dramatic.
- Typography is system-native and editorially calm.

## Colors

The palette is built from paper whites, graphite neutrals, and one near-black action tone. Color is used structurally, not decoratively.

### Primary
- **Graphite Action** (`#222325`): Primary buttons, active navigation states, and high-importance actions. This is the darkest recurring interactive tone, but it should still read as a neutral rather than a brand color.

### Neutral
- **Ink Body** (`#161719`): Main body text, major headings, and the most legible foreground color.
- **Deep Ink** (`#111214`): Used sparingly for stronger emphasis and darker hover states.
- **Slate Utility** (`#5F6368`): Secondary copy, supporting labels, field descriptions, and calm UI chrome.
- **Mist Label** (`#80858B`): Eyebrows, tertiary labels, and less important metadata.
- **Canvas** (`#F2F2F0`): The base page background. This is the main field of the app and should feel like warm paper rather than clean white.
- **Elevated Canvas** (`#F7F7F5`): Secondary background for previews, inset surfaces, and lighter protected areas.
- **Surface** (`#FBFBF9`): Cards, forms, tables, and floating panels.
- **Muted Surface** (`#EBEBE8`): Divider-adjacent background and quiet layering when a second neutral plane is needed.
- **Subtle Border** (`#1112141A`): Default dividers and container strokes.
- **Strong Border** (`#1112142E`): Hover emphasis, stronger separations, and rare high-contrast boundary moments.

### Named Rules
**The Monochrome Discipline Rule.** Accent is still neutral. If a screen can be made more expressive by adding blue, green, or purple, the design should first solve the problem with hierarchy, spacing, and contrast.

**The Soft Contrast Rule.** Never use pure `#000000` or `#FFFFFF`. Every black and white in the system is slightly softened to preserve a print-like calmness.

## Typography

**Display Font:** system sans stack with Chinese support (`-apple-system`, `"Segoe UI"`, `"PingFang SC"`, `"Hiragino Sans GB"`, `"Microsoft YaHei"`, sans-serif)
**Body Font:** system sans stack with Chinese support (`-apple-system`, `"Segoe UI"`, `"PingFang SC"`, `"Hiragino Sans GB"`, `"Microsoft YaHei"`, sans-serif)
**Label/Mono Font:** no distinct mono family in the current system

**Character:** Typography is native, quiet, and efficient. It should feel familiar on macOS and Windows, with enough weight contrast to create order, but never enough stylization to distract from editorial tasks.

### Hierarchy
- **Display** (600, 42px, 1.05): Login hero and rare large framing moments only.
- **Headline** (600, 31px, 1.1): Primary page titles at the top of each admin route.
- **Title** (600, 20px, 1.2): Section headings inside cards, dashboards, and inline editors.
- **Body** (400, 14px, 1.6): Standard interface copy, helper text, and table-adjacent prose. Long-form descriptive lines should stay around 65 to 75 characters where practical.
- **Label** (600, 11px, 0.08em): Eyebrows, overlines, and compact metadata labels in uppercase.

### Named Rules
**The One Family Rule.** This interface uses one system-native sans family for nearly everything. Do not introduce display faces, serif accents, or decorative label fonts.

**The Weight Before Size Rule.** Create hierarchy through weight and spacing before reaching for much larger type.

## Elevation

Elevation is ambient and structural. Most surfaces are flat in spirit, but lightly lifted through borders and soft shadows so the admin never feels raw or unfinished. The page itself sits on a paper-toned field, while cards and panels float just enough to separate work areas.

### Shadow Vocabulary
- **Panel Lift** (`box-shadow: 0 16px 40px rgba(17, 18, 20, 0.06)`): General cards, forms, and quiet elevated surfaces.
- **Strong Panel Lift** (`box-shadow: 0 20px 56px rgba(17, 18, 20, 0.08)`): Higher-emphasis overlays or denser foreground panels that need a bit more authority.
- **Shell Lift** (`box-shadow: 0 18px 60px rgba(17, 18, 20, 0.08)`): Large structural surfaces such as the left navigation shell and major content containers.

### Named Rules
**The Flat-By-Default Rule.** Components should feel flat until they need separation. Shadows exist to distinguish planes, not to decorate every element.

**The Bottom-Safe Rule.** Any scroll container must preserve room for bottom radius and shadow. Do not clip card shadows against layout edges.

## Components

Each component family should feel restrained, efficient, and editorially calm. Rounded corners and shadows are present, but always secondary to clarity.

### Buttons
- **Shape:** restrained rounded rectangle (`10px` default global radius)
- **Primary:** graphite fill with surface text, compact horizontal padding, no loud accent colors
- **Hover / Focus:** slightly darker neutral shift, not a hue shift
- **Secondary / Ghost / Tertiary:** built from surface backgrounds and neutral foregrounds, never from tinted brand colors

### Cards / Containers
- **Corner Style:** medium-large rounding (`18px` on major surfaces, `14px` on sub-surfaces)
- **Background:** warm white surfaces on a slightly darker paper canvas
- **Shadow Strategy:** ambient lift only, using the panel or shell shadow tokens
- **Border:** always present as a subtle ink border
- **Internal Padding:** most major cards use `24px`, denser inline editing surfaces can step down to `18px`

### Inputs / Fields
- **Style:** white surface, neutral text, restrained radius, no colorized fills
- **Focus:** controlled contrast increase rather than saturated outline glow
- **Error / Disabled:** errors should stay muted and structural; disabled fields should lower contrast but remain legible

### Navigation
- **Style:** the left navigation lives inside its own framed shell with rounded outer corners and quiet shadow
- **Typography:** medium-weight system sans, no icon-driven exaggeration
- **Default / Hover / Active:** default is neutral, hover is subtle, active state flips to dark neutral background with light text
- **Mobile Treatment:** if collapse returns in the future, it must become a real icon-first navigation state, not a squeezed empty strip

### Tags and Status Chips
- **Style:** low-contrast pills used for state and metadata, not as bright badges
- **State:** status chips may darken or lighten in neutral space, but should not introduce a new palette family

## Do's and Don'ts

### Do:
- **Do** keep the interface monochrome. Reach for `#222325`, `#161719`, and the paper whites before adding any new accent.
- **Do** use `18px` as the upper end for major surface rounding and `14px` or below for nested surfaces.
- **Do** preserve the paper hierarchy: canvas background first, white surface second, border third, shadow last.
- **Do** use uppercase 11px labels for overlines and section eyebrows only.
- **Do** let spacing create elegance. The system gets its refinement from rhythm and restraint, not ornament.

### Don't:
- **Don't** use pure `#000000` or `#FFFFFF`.
- **Don't** reintroduce saturated theme colors such as green, blue, purple, or neon accents into buttons, links, or navigation.
- **Don't** push corner radius back into the oversized range. Anything beyond `22px` on standard product surfaces will feel too soft for this system.
- **Don't** use gradient text, glossy fills, glassmorphism, or decorative blur-heavy cards.
- **Don't** collapse navigation into an empty strip. If a collapsed sidebar returns, it must remain fully legible and intentionally designed.
